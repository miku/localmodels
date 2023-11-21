.PHONY:
create-custom-models:
	ollama create llama-mario -f custom/Modelfile.mario
	# ollama run llama-mario
