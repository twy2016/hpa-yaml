yum install httpd-tools

ab -c 100 -n 1000000000 http://$(kubectl get po -o wide| grep web-0| awk '{print $6}'):80/