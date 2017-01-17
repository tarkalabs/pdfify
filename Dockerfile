FROM vagmi/wkhtmltopdf:latest
EXPOSE 8300
ADD pdfify pdfify
ENTRYPOINT ["/pdfify"]
