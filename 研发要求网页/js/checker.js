const form = document.getElementById("myForm")

function validateForm(){
    const emailID = document.getElementById("email")
    const pwdID = document.getElementById("password")

    const email = pwdID.value
    const password = pwdID.value
    
    if(email == "" || password == ""){
        alert("所有输入框必填")
        emailID.style.borderColor = "red"
        pwdID.style.borderColor = "red"
    }else{
    if(email !== "admin@qq.com" && password !== "admin"){
        alert("请检查你的用户名或密码")
        emailID.style.borderColor = "red"
        pwdID.style.borderColor = "red"
      }else{
        alert("成功")
        emailID.style.borderColor = "black"
        pwdID.style.borderColor = "black"
    }
  }
};