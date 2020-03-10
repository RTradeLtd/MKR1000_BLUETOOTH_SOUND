# Examples

https://gist.github.com/dmiddlecamp/ff3d08615110618a3046
https://www.hackster.io/middleca/sending-sound-over-the-internet-f097b4
https://forums.adafruit.com/viewtopic.php?f=19&t=54023&p=273033&hilit=1063#p273033


```

var outStream = fs.createWriteStream("test.wav");

var writeHeader = function() {
	var b = new Buffer(1024);
	b.write('RIFF', 0);
	/* file length */
	b.writeUInt32LE(32 + samplesLength * 2, 4);
	//b.writeUint32LE(0, 4);

	b.write('WAVE', 8);
	/* format chunk identifier */
	b.write('fmt ', 12);

	/* format chunk length */
	b.writeUInt32LE(16, 16);

	/* sample format (raw) */
	b.writeUInt16LE(1, 20);

	/* channel count */
	b.writeUInt16LE(1, 22);

	/* sample rate */
	b.writeUInt32LE(sampleRate, 24);

	/* byte rate (sample rate * block align) */
	b.writeUInt32LE(sampleRate * 2, 28);

	/* block align (channel count * bytes per sample) */
	b.writeUInt16LE(2, 32);

	/* bits per sample */
	b.writeUInt16LE(16, 34);

	/* data chunk identifier */
	b.write('data', 36);

	/* data chunk length */
	//b.writeUInt32LE(40, samplesLength * 2);
	b.writeUInt32LE(0, 40);


	outStream.write(b.slice(0, 50));
};


```