FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY k8-heart-beat /
CMD ["/k8-heart-beat"]