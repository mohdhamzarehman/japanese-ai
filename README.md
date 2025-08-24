 # Japanese Text-to-Speech CLI
>> 
>> A command-line tool for converting Japanese text to speech using Google Cloud Text-to-Speech API. This tool supports multiple voice types and audio formats, making it easy to generate natural-sounding Japanese speech.
>> 
>> ## Features
>> 
>> - Multiple voice types support (Standard and Wavenet voices)
>> - Adjustable speech rate and pitch
>> - Multiple output formats (WAV and MP3)
>> - Easy-to-use command-line interface
>> - Built with Google Cloud Text-to-Speech API
>> 
>> ## Prerequisites
>> 
>> - Go 1.24.1 or later
>> - Google Cloud Platform account
>> - Service account key with Text-to-Speech API access
>> - Set up GOOGLE_APPLICATION_CREDENTIALS environment variable
>> 
>> ## Installation
>> 
>> ```bash
>> go get github.com/mohdhamzarehman/tt-japanese
>> ```
>> 
>> ## Usage
>> 
>> Basic usage:
>> 
>> ```bash
>> tt-japanese -text 'こんにちは世界' -o output.mp3
>> ```
>> 
>> Advanced options:
>> 
>> ```bash
>> tt-japanese -text 'こんにちは世界' -voice 'wavenet-a' -rate 1.2 -pitch 2.0 -o output.mp3
>> ```
>> 
>> ### Parameters
>> 
>> - `-text`: Text to convert to speech (required)
>> - `-voice`: Voice type to use (default: 'stand-a')
>>   - Available voices: stand-a, stand-b, stand-c, stand-d, wavenet-a, wavenet-b, wavenet-c, wavenet-d
>> - `-rate`: Speech rate (0.25 to 4.0, default: 1.0)
>> - `-pitch`: Speaking pitch (-20.0 to 20.0, default: 0.0)
>> - `-o`: Output file path (supported formats: .wav, .mp3)
>> 
>> ## Voice Types
>> 
>> ### Standard Voices
>> - stand-a: Standard voice type A
>> - stand-b: Standard voice type B
>> - stand-c: Standard voice type C
>> - stand-d: Standard voice type D
>> 
>> ### Wavenet Voices (Higher Quality)
>> - wavenet-a: WaveNet voice type A
>> - wavenet-b: WaveNet voice type B
>> - wavenet-c: WaveNet voice type C
>> - wavenet-d: WaveNet voice type D
>> 
>> ## License
>> 
>> This project is licensed under the MIT License.
>> 
>> ## Contributing
>> 
>> Contributions are welcome! Please feel free to submit a Pull Request.
>> " > README.md }
>> 
