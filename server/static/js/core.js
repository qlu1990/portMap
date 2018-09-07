var data = {"es":{"control":"启动","info":"停止","control-id":"es-start"},
              "mysql":{"control":"启动","info":"停止","control-id":"mysql-start"},
              "hadoop":{"control":"启动","info":"停止","control-id":"hadoop-start"},
              "hbase":{"control":"启动","info":"停止","control-id":"hbase-start"},
              "cas":{"control":"启动","info":"停止","control-id":"cas-start"},
              "neo4j":{"control":"启动","info":"停止","control-id":"neo4j-start"},
              "backend":{"control":"启动","info":"停止","control-id":"backend-start"},
              "map":{"control":"启动","info":"停止","control-id":"map-start"},
              "front":{"control":"启动","info":"停止","control-id":"front-start"}
             }
  var app = new Vue({
    el: '#app',
    data: data,
    methods:{
        changedata: function(jsondata){

            for(var item in jsondata){
                this.data[item]["info"] = jsondata[item].info;
                if(jsondata[item].status == "started")
                {
                    this.data[item]["control"] = "停止";
                }
                else {
                    this.data[item]["control"] = "启动";
                }
            }

        }
        
    }
  })

$(function(){
    getStatus()
    setInterval(getStatus,5000);
//getLoc();
 });
 function getStatus(){
    $.get("/status", function(jsondata,status){
        for(var item in jsondata){
            data[item]["info"] = jsondata[item].info;
            if(jsondata[item].status == "started")
            {
                data[item]["control"] = "停止";
                data[item]["control-id"] = item+"-stop";
            }
            else {
                data[item]["control"] = "启动";
                data[item]["control-id"] = item+"-start";
            }
        }
    }
    );
 }
 function setStatus(data){
     $.each(data,function(index,value)
    {
        alert(index.name);
        // console.log(index.name);
        // console.log(value);
    });
 }
$("button").click(function(){

    var control = $(this).text();
    var operate = "";
    if(control == "启动")
    {
        operate = "start";
    }else if(control == "停止"){
        operate = "stop";
    }else{
        operate = "start";
    }
    var softname = $(this).attr("id");
    var str_start = softname.indexOf("all");
    if(str_start == 0)
    {
        softname = "all";
    }
    $.post("/control",
    {
      "softname":softname,
      "operate":operate
    },
    );
  });
  

  