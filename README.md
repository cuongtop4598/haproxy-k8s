# HA proxy

## Run model
### step 1: 
> minikube start --driver=kvm2
### step 2:
> kubectl apply -f .

## Test model
### step 1: Chạy test script gửi request đến nginx-app theo tần số 100ms
> go run ./test/main.go
### step 2: Giả lập pod nginx-app down/up
> kubectl delete pod --all
### step 3: Kiểm tra log test script trong terminal
> Ta thấy khi pod down hết sẽ không tạo được request, nhưng khi 1 trong 3 pod down thì không ảnh hưởng tới request.