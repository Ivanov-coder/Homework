const form = document.getElementById("myForm")
const pwdRex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;  

function validateForm(){
    const emailID = document.getElementById("email")
    const pwdID = document.getElementById("password")
    const ckpwdID = document.getElementById("checkPassword")

    const email = emailID.value
    const password = pwdID.value
    const confirmPassword = ckpwdID.value

    if(email == "" || password == "" || confirmPassword == ""){
        alert("所有输入框必填")
        emailID.style.borderColor = "red"
        pwdID.style.borderColor = "red"
        ckpwdID.style.borderColor = "red"
    }

    if(password !== "" && confirmPassword !== "" && password !== confirmPassword){
        alert("两次输入密码不一致")
        pwdID.style.borderColor = "red"
        ckpwdID.style.borderColor = "red"
    }else{

    if(email !== "" && password.length >= 8 && confirmPassword !== "" ){
        if (pwdRex.test(password)) {  
            alert("成功")
            emailID.style.borderColor = "black"
            pwdID.style.borderColor = "black"
            ckpwdID.style.borderColor = "black"
            emailID.value = ""
            pwdID.value = ""
            ckpwdID.value = "" 
          } else {  
            alert("无效的密码\n请确保密码含有大小写英文字符和特殊符号@$!%*?&");  
            pwdID.style.borderColor = "red"
            ckpwdID.style.borderColor = "red"
          }
    }else if(email !== "" && password.length < 8 && confirmPassword !== ""){
        alert("密码长度必须大于8")
        pwdID.style.borderColor = "red"
        ckpwdID.style.borderColor = "red"
    }
    }
}