from fastapi import FastAPI, UploadFile, File
from fastapi.responses import FileResponse
from starlette.responses import Response
import os
from prometheus_client import Counter, generate_latest, CONTENT_TYPE_LATEST

app = FastAPI()

# Metrics
upload_counter = Counter("image_upload_total", "Total uploaded images")

# Create upload directory if not exists
os.makedirs("uploads", exist_ok=True)

@app.get("/healthz")
def healthz():
    return "ok"

@app.get("/ready")
def ready():
    return "ready"

@app.post("/upload")
async def upload_image(file: UploadFile = File(...)):
    upload_counter.inc()
    file_location = f"uploads/{file.filename}"
    with open(file_location, "wb") as f:
        f.write(await file.read())
    return {"status": "uploaded ✅", "file": file.filename}

@app.get("/images/{filename}")
def get_image(filename: str):
    file_path = f"uploads/{filename}"
    return FileResponse(file_path)

@app.get("/metrics")
def metrics():
    data = generate_latest()
    return Response(content=data, media_type=CONTENT_TYPE_LATEST)
