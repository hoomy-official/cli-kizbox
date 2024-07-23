FROM scratch

COPY cli-kizbox /usr/bin/cli-kizbox

ENTRYPOINT [ "/usr/bin/cli-kizbox" ]

CMD ["serve"]

