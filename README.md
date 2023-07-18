# k8s-job-operator
帮助程序员在k8s上运行特定业务job，可随时执行，并提供简单web页面供产品和运营等非开发同事使用。

## 使用
### 1. 配置
   - 首先需要先配置kubeconfig-local文件，这个是kubectl的登陆配置，否则无法使用client-go operator来操作k8s集群。
   - 除了go 1.20的环境，你还需要配置npm，因为这个项目结合来react的前端页面，所以需要npm来编译前端页面。
   - 配置好后，执行`make`命令，会自动编译出可执行文件`k8s-job-operator`，以及前端页面的静态文件。
## 截图
![image](https://github.com/zealerFT/k8s-job-operator/blob/main/resource/before.png)
![image](https://github.com/zealerFT/k8s-job-operator/blob/main/resource/done.png)

## 结语
整个项目其实是一个k8s的operator，但是这个operator的功能比较简单，只是为了帮助程序员在k8s上运行特定业务job，可随时执行，并提供简单web页面供产品和运营等非开发同事使用。
如果你看了代码加会发现client-go写的很好，go的云原生支持确实好，你自定义自己的功能。