class Ctw < Formula
  desc "ctw - A terminal based app for watching cryptocurrency rate in different currency"
  homepage "https://github.com/eftakhairul/crypto-terminal-watch"
  url "http://apps.eftakhairul.com/ctw.tar.gz"
  version "0.0.1"
  sha256 "4f6376cfcb4c8ea464aa39f3808c7ee3bb0bf23d301ea369582a1fd6809105f1"

  def install
    bin.install "ctw"
  end
end