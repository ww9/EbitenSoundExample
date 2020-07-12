# Ebiten Sound Example

Small repository to teach myself to play sounds in Ebiten game engine (https://ebiten.org/).

## wav: bits per sample must be 8 or 16 but was 24

If you get this error while loading a WAV file, you can convert the wav file to 16 bits using ffmpeg commandline utility like this:

`ffmpeg -i file.wav -acodec pcm_s16le -ac 1 -ar 48000 file.wav`

You can also use an online free service like: https://audio.online-convert.com/convert-to-wav

Make sure to set `Change bit resolution` to 16 bits on that website.

![Screenshot](/screenshot.png)