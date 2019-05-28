# myConverter

## Features
- This application can encrypt / decrypt files by using pass-phrase encryption.  
- Encryption with AES CTR mode.

## Note:
- ディレクトリの再帰的な探索は path/filepath.Walk を使うのが一番近道だと思った。  
    - しかし、filepath.Walk 内で呼び出せる関数が WalkFunc に限定されているようで、かつ WalkFunc の型が決まっていたので少しやり辛かった。
        - これは関数ポインタを渡したり、無名関数を使うような仕様の方が良いのでは・・・
    - 一方で、これは[先月の Software Design](https://gihyo.jp/magazine/SD/archive/2019/201905) で見た Generator Pattern を試すと Go らしくなるのでは考えた。
    - [同じことを考える人](https://gist.github.com/sethamclean/9475737)がいたので、参考にした。
