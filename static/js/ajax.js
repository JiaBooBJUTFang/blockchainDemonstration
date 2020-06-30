function ajaxget(url,fnSucc,fnField)
{
    if(window.XMLHttpRequest)
        {
            var oAjax = new XMLHttpRequest();
        }
        else
        {
            var oAjax = new ActiveXObject("Microsoft.XMLHTTP");//IE6浏览器创建ajax对象
        }
        oAjax.open("GET",url,true);//把要读取的参数的传过来。
        oAjax.send();
        oAjax.onreadystatechange=function()
        {
            if(oAjax.readyState==4)
            {
                if(oAjax.status==200)
                {
                    fnSucc(oAjax.responseText);//成功的时候调用这个方法
                }
                else
                {
                    if(fnfiled)
                    {
                        fnField(oAjax.status);
                    }
                }
            }
        };
}
function ajaxpost(url,obj,fnSucc1){
var httpRequest = new XMLHttpRequest();//第一步：创建需要的对象
httpRequest.open('POST', url, true); //第二步：打开连接
httpRequest.setRequestHeader("Content-type","application/x-www-form-urlencoded");//设置请求头 注：post方式必须设置请求头（在建立连接后设置请求头）
//httpRequest.send('name=teswe&ee=ef');//发送请求 将写在send中
httpRequest.send(JSON.stringify(obj));
/**
 * 获取数据后的处理程序
 */
httpRequest.onreadystatechange = function () {//请求后的回调接口，可将请求成功后要执行的程序写在其中
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {//验证请求是否发送成功
        // var json = httpRequest.responseText;//获取到服务端返回的数据
        // console.log(json);
        fnSucc1(httpRequest.responseText);
    }
};
}
function ajaxjson(url,obj,fnSucc2){
var httpRequest = new XMLHttpRequest();//第一步：创建需要的对象
httpRequest.open('POST', url, true); //第二步：打开连接/***发送json格式文件必须设置请求头 ；如下 - */
httpRequest.setRequestHeader("Content-type","application/json");//设置请求头 注：post方式必须设置请求头（在建立连接后设置请求头）var obj = { name: 'zhansgan', age: 18 };
httpRequest.send(JSON.stringify(obj));//发送请求 将json写入send中
/**
 * 获取数据后的处理程序
 */
httpRequest.onreadystatechange = function () {//请求后的回调接口，可将请求成功后要执行的程序写在其中
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {//验证请求是否发送成功
        var json = httpRequest.responseText;//获取到服务端返回的数据
        console.log(json);
        fnSucc2(httpRequest.responseText);
    }
};
}