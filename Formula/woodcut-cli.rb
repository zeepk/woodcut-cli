# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class WoodcutCli < Formula
  desc "A simple description of your application."
  homepage "https://github.com/zeepk/woodcut-cli"
  version "0.1.9"

  on_macos do
    url "https://github.com/zeepk/woodcut-cli/releases/download/v0.1.9/woodcut-cli_0.1.9_darwin_amd64.tar.gz"
    sha256 "0cf715d435042ba3daf79ed2a03af18301e26c80232c0144c139615300233c14"

    def install
      bin.install "woodcut"
    end

    if Hardware::CPU.arm?
      def caveats
        <<~EOS
          The darwin_arm64 architecture is not supported for the WoodcutCli
          formula at this time. The darwin_amd64 binary may work in compatibility
          mode, but it might not be fully supported.
        EOS
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/zeepk/woodcut-cli/releases/download/v0.1.9/woodcut-cli_0.1.9_linux_amd64.tar.gz"
      sha256 "48a1d124b8337acda591e7202489b7639364401d5b367c2e91fc498093f82500"

      def install
        bin.install "woodcut"
      end
    end
  end
end
