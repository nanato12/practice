"use strict";
var User = /** @class */ (function () {
    function User(familyName, givenName, age) {
        this.age = age;
        this.familyName = familyName;
        this.givenName = givenName;
    }
    return User;
}());
var user = new User('海老原', '賢次', 44); // 名前と年齢は適当に
var contentsElem = document.getElementById('contents');
if (!!contentsElem) {
    contentsElem.innerText = user.familyName + " " + user.givenName;
}
