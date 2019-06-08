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

module MRuby
  class Build
    # Enable compilation of zlib.
    #
    # @return [ Void ]
    def enable_zlib
      [cc, cxx].each do |cc|
        cc.defines += %w[LIBSSH2_HAVE_ZLIB ZLIB_STATIC HAVE_UNISTD_H]
      end
    end

    # Enable linking against openssl instead of mbedtls.
    #
    # @return [ Void ]
    def enable_openssl
      linker.libraries << 'crypto'

      [cc, cxx].each do |cc|
        cc.defines += %w[MRB_SSH_LINK_CRYPTO LIBSSH2_OPENSSL]
      end
    end

    # Enable compiler optimizations.
    #
    # @return [ Void ]
    def enable_optimizations
      [cc, cxx].each do |cc|
        cc.flags << (toolchains.include?('clang') ? '-Oz' : '-Os')
      end
    end

    # Enable static build instead of linking with shared libs.
    #
    # @return [ Void ]
    def enable_static
      linker.flags_before_libraries << '-Wl,-Bstatic'
    end

    # Enable threading mode for mbedtls.
    #
    # @return [ Void ]
    def enable_mbedtls_threading
      [cc, cxx].each do |cc|
        cc.defines += %w[MBEDTLS_THREADING_PTHREAD MBEDTLS_THREADING_C]
      end
    end

    # Set the proper include headers to ensure that the binary
    # wont depend on newer versions of glibc.
    #
    # param [ String ] version The maximun supported glibc version.
    #
    # @return [ Void ]
    def glibc_version=(version)
      return if !ENV['GLIBC_HEADERS'] || is_a?(MRuby::CrossBuild)

      [cc, cxx].each do |cc|
        cc.flags << "-include #{ENV['GLIBC_HEADERS']}/x64/force_link_glibc_#{version}.h"
      end
    end
  end
end
