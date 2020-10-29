


If you need to have imagemagick@6 first in your PATH run:
  echo 'export PATH="/usr/local/opt/imagemagick@6/bin:$PATH"' >> ~/.bash_profile

#For compilers to find imagemagick@6 you may need to set:
export LDFLAGS="-L/usr/local/opt/imagemagick@6/lib"
export CPPFLAGS="-I/usr/local/opt/imagemagick@6/include"

#For pkg-config to find imagemagick@6 you may need to set:
export PKG_CONFIG_PATH="/usr/local/opt/imagemagick@6/lib/pkgconfig"

export CGO_CFLAGS_ALLOW='-Xpreprocessor'

pkg-config --cflags --libs MagickWand


go get -u gopkg.in/gographics/imagick.v2/imagick





 gcc \
   -I /usr/local/Cellar/imagemagick@6/6.9.10-27/include/ImageMagick-6  \
   -L /usr/local/Cellar/imagemagick@6/6.9.10-27/lib \
   resize.c

pkg-config --cflags --libs MagickWand

-DMAGICKCORE_HDRI_ENABLE=0 -DMAGICKCORE_QUANTUM_DEPTH=16 -DMAGICKCORE_HDRI_ENABLE=0 -DMAGICKCORE_QUANTUM_DEPTH=16 -I/usr/local/Cellar/imagemagick@6/6.9.10-27/include/ImageMagick-6 -L/usr/local/Cellar/imagemagick@6/6.9.10-27/lib -lMagickWand-6.Q16 -lMagickCore-6.Q16



 c99 -DMAGICKCORE_HDRI_ENABLE=0 -DMAGICKCORE_QUANTUM_DEPTH=16 -DMAGICKCORE_HDRI_ENABLE=0 -DMAGICKCORE_QUANTUM_DEPTH=16 -I/usr/local/Cellar/imagemagick@6/6.9.10-27/include/ImageMagick-6 -L/usr/local/Cellar/imagemagick@6/6.9.10-27/lib -lMagickWand-6.Q16 -lMagickCore-6.Q16 resize.c