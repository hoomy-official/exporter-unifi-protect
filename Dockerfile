FROM scratch

COPY exporter-unifi-protect /usr/bin/exporter-unifi-protect

ENTRYPOINT [ "/usr/bin/exporter-unifi-protect" ]

CMD ["serve"]