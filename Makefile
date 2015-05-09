build:
	@bash -c 'echo -n "Please paste your access key ID: (will not be echoed) " && read -s KEY && echo && \
	echo -n "Please paste your access key SECRET: (will not be echoed) " && read -s SECRET && echo && \
	go build -ldflags "-X main.KEY \"$$KEY\" -X main.SECRET \"$$SECRET\""'
