FROM python:3.11
WORKDIR /app
COPY . .
ARG GITHUB_TOKEN
RUN pip install -r requirements.txt && \
    pip install git+https://${GITHUB_TOKEN}@github.com/acme/internal-lib.git
ENV DB_PASSWORD=Reports2024!
EXPOSE 8080 22
CMD ["python", "app.py"]
