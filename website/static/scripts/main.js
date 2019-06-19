const host = "http://localhost:7111/";

window.config = {
    http: '111111'
};

window.urlConf = {
    getAllMenu: host + "menu/get",
    lookMenu: host + "menu/user/look",
    insertMenu: host + "menu/inert",
    deleteMenu: host + "menu/delete",
    alterMenu: host + "menu/alter",
    randomMenu: host + "menu/random",
};


(function ($) {
    $.getUrlParam = function (name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)")
        var r = window.location.search.substr(1).match(reg)
        if (r != null) return unescape(r[2])
        return null
    }
})(jQuery)

/*
* 方法:Array.remove(dx) 通过遍历,重构数组
* 功能:删除数组元素.
* 参数:dx删除元素的下标.
*/
Array.prototype.remove = function (dx) {
    if (isNaN(dx) || dx > this.length) {
        return false;
    }
    for (var i = 0, n = 0; i < this.length; i++) {
        if (this[i] != this[dx]) {
            this[n++] = this[i]
        }
    }
    this.length -= 1
}

