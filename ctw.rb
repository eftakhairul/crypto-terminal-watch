class Ctw < Formula
  desc "ctw - A terminal based app for watching cryptocurrency rate in different currency"
  homepage "https://github.com/eftakhairul/crypto-terminal-watch"
  url "http://apps.eftakhairul.com/ctw.tar.gz"
  version "0.0.1"
  sha256 "f9e808910fdecfba9a07467320c199c69e711ff23bd078d9fb646d8d5c33ae74"  

  bottle :unneeded

  def install
    bin.install "ctw"
  end
end