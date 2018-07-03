# Apache 2.0 License
#
# Copyright (c) 2018 Sebastian Katzer, appPlant GmbH
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

module SKI
  # Colorized ORBIT logo
  LOGO = <<-LOGO.chomp.set_color(OS.posix? ? 208 : :light_yellow).freeze
          `-/++++++/:-`
      `/yhyo/:....-:/+syyy+:`
    .yd+`                `-+yds:     `
   om/                        -+``sdyshh.
  sd`                            hy    om
 /N.                             od.  `yd
 do                               /yyyy+`.`
`N:                                      /ms`
`M-                                        +m+
 m+                                         .+`
 od                                            +`
 .N:                     -oyyyyyo-             N.         +m.   /.
  om`                  -dy-     -yd-    `  .`  N. `.`           d/``
   hy                 -N-         -N-   Nss+/  Nys+/oh+   -m  -oNyoo
   `dy                yy           yy   N/     N-    `m:  -m    d:
    `dy               om           mo   N.     N.     h/  -m    d:
      yd.              hd.       .hh`   N.     N.    -N.  -m    d:
       +m:              :hho///ohh:     N.     Nys++sy-   -m    oh+o
        .dy`               .:::.                  ``              ``
          +m+
           `od+
             `od+`
               `+ds
LOGO
end
