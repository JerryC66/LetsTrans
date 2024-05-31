from abc import ABC, abstractmethod
import json
import grpc
from concurrent import futures
import document_processor_pb2
import document_processor_pb2_grpc


class DocumentProcessor(ABC):
    @abstractmethod
    def read_document(self, content: str) -> list:
        pass

    @abstractmethod
    def modify_document(self, content: str, replacements: list) -> str:
        pass


class TxtDocumentProcessor(DocumentProcessor):
    def read_document(self, content: str) -> list:
        segments = []
        segment_id = 0
        for line in content.splitlines():
            if line.strip() == "":
                continue
            segments.append({
                "id": segment_id,
                "original_text": line.strip()
            })
            segment_id += 1
        return segments

    def modify_document(self, content: str, replacements: list) -> str:
        lines = []
        segment_id = 0
        for line in content.splitlines():
            if line.strip() == "":
                lines.append(line)
                continue
            modified_line = replacements[segment_id].get('translated_text', line.strip())
            lines.append(modified_line)
            segment_id += 1
        return "\n".join(lines)


class DocxDocumentProcessor(DocumentProcessor):
    def read_document(self, content: str) -> list:
        from docx import Document
        from io import BytesIO
        doc = Document(BytesIO(content.encode('utf-8')))
        segments = []
        segment_id = 0
        for paragraph in doc.paragraphs:
            if paragraph.text.strip() == "":
                continue
            segments.append({
                "id": segment_id,
                "original_text": paragraph.text.strip()
            })
            segment_id += 1
        return segments

    def modify_document(self, content: str, replacements: list) -> str:
        from docx import Document
        from io import BytesIO
        doc = Document(BytesIO(content.encode('utf-8')))
        segment_id = 0
        for paragraph in doc.paragraphs:
            if paragraph.text.strip() == "":
                continue
            paragraph.text = replacements[segment_id]['translated_text']
            segment_id += 1
        output = BytesIO()
        doc.save(output)
        return output.getvalue().decode('utf-8')


class PptxDocumentProcessor(DocumentProcessor):
    def read_document(self, content: str) -> list:
        from pptx import Presentation
        from io import BytesIO
        prs = Presentation(BytesIO(content.encode('utf-8')))
        segments = []
        segment_id = 0
        for slide in prs.slides:
            for shape in slide.shapes:
                if not shape.has_text_frame:
                    continue
                text_frame = shape.text_frame
                original_text = text_frame.text
                if original_text.strip() == "":
                    continue
                segments.append({
                    "id": segment_id,
                    "original_text": original_text.strip()
                })
                segment_id += 1
        return segments

    def modify_document(self, content: str, replacements: list) -> str:
        from pptx import Presentation
        from io import BytesIO
        prs = Presentation(BytesIO(content.encode('utf-8')))
        segment_id = 0
        for slide in prs.slides:
            for shape in slide.shapes:
                if not shape.has_text_frame:
                    continue
                text_frame = shape.text_frame
                if text_frame.text.strip() == "":
                    continue
                text_frame.text = replacements[segment_id]['translated_text']
                segment_id += 1
        output = BytesIO()
        prs.save(output)
        return output.getvalue().decode('utf-8')


class PdfDocumentProcessor(DocumentProcessor):
    def read_document(self, content: str) -> list:
        import fitz  # PyMuPDF
        from io import BytesIO
        doc = fitz.open("pdf", BytesIO(content.encode('utf-8')))
        segments = []
        segment_id = 0
        for page_num in range(len(doc)):
            page = doc[page_num]
            blocks = page.get_text("blocks")
            for block in blocks:
                text = block[4]
                if text.strip() == "":
                    continue
                segments.append({
                    "id": segment_id,
                    "original_text": text.strip()
                })
                segment_id += 1
        return segments

    def modify_document(self, content: str, replacements: list) -> str:
        lines = '\n'.join([item['translated_text'] for item in replacements])
        return lines


# gRPC 服务实现

class DocumentProcessorServicer(document_processor_pb2_grpc.DocumentProcessorServicer):
    def __init__(self):
        self.processors = {
            'txt': TxtDocumentProcessor(),
            'docx': DocxDocumentProcessor(),
            'pptx': PptxDocumentProcessor(),
            'pdf': PdfDocumentProcessor()
        }

    def ProcessDocument(self, request, context):
        processor = self.processors.get(request.file_type)
        if not processor:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details('Unsupported file type')
            return document_processor_pb2.ProcessDocumentResponse()

        segments = processor.read_document(request.content)
        return document_processor_pb2.ProcessDocumentResponse(segments=json.dumps(segments))

    def ModifyDocument(self, request, context):
        processor = self.processors.get(request.file_type)
        if not processor:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details('Unsupported file type')
            return document_processor_pb2.ModifyDocumentResponse()

        replacements = json.loads(request.replacements)
        modified_content = processor.modify_document(request.content, replacements)
        return document_processor_pb2.ModifyDocumentResponse(modified_content=modified_content)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    document_processor_pb2_grpc.add_DocumentProcessorServicer_to_server(DocumentProcessorServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
