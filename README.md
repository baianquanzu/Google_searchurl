# Google_searchurl
用于爬取谷歌关键词搜索的url，便于红队，src等快速提取</br>
使用方式：</br>
  源码直接运行需要解决：</br>
    go的环境，当出现下面的报错</br>
    go: go.mod file not found in current directory or any parent directory; see 'go help modules'</br>
    运行：</br>
    go env -w GO111MODULE=on</br>
    go mod init xxx //xxx代表文件名</br>
    可以直接编译：go build -o crawl_urls.exe url.go
  </br>

直接使用exe文件：</br>
直接找到文件存储目录运行cmd输入：Google_searchurl.exe </br>
![image](https://github.com/baianquanzu/Google_searchurl/assets/47970894/18a09c4f-0ad4-460e-b853-cbf6f2d7dc74)  </br>
这里可以设置你的爬取数量和代理，这里代理默认是http的，也可以设置socks5</br>
