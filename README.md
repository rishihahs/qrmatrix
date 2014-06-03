qrmatrix
=========

qrmatrix encodes large data by using structured-appended symbols (thereby splitting it into multiple QR Codes)

Dependencies
-----------

[libqrencode] is required.

You can build it yourself, or use a package manager like homebrew/apt-get/etc:

    brew install libqrencode

Installation
--------------

    go get github.com/rishihahs/qrmatrix
    
Usage
------

    Usage of qrmatrix:
      --codes-per-row: [optional] number of qr codes per row
      --output: [optional] The file to output the image to (e.g. codes.png). Will output to STDOUT by default
      --size: [optional] width and height (in pixels) of each module (square block of qr code)

    Data sent through STDIN will be encoded.
    E.g.
	    qrmatrix --size 3 < fileToEncode
	    echo "Hello World" | qrmatrix > output.png

Example
----------

    curl -s 'http://zeptojs.com/zepto.min.js' | qrmatrix > out.png
    
![Zepto.js QR Code](http://i.imgur.com/qKceZci.png)

License
----

MIT

[libqrencode]:http://fukuchi.org/works/qrencode/index.html.en

