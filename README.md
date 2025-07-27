Release script:
./ci-script.sh

budowa dockera:

docker buildx build -t translate_server_image -f dockerfile .

run tłumacza:

echo "Jack would like buy a new car ." > input.txt
docker run --gpus all -v C:\Users\K\Desktop\TS:/data translate_server_image:latest marian-decoder -m /opt/marian-dev/vocab/opus+bt.spm32k-spm32k.transformer-align.model1.npz.best-perplexity.npz -v /opt/marian-dev/vocab/opus+bt.spm32k-spm32k.vocab.yml /opt/marian-dev/vocab/opus+bt.spm32k-spm32k.vocab.yml -i /data/input.txt -o /data/output.txt

run servera:
go build cmd/TranslateServer/main.go
docker run -p 5000:5000 -v C:\Users\K\Desktop\TS:/data translate_server_image:latest /data/main
