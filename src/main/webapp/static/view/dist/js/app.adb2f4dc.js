(function(e){function t(t){for(var r,a,o=t[0],i=t[1],s=t[2],l=0,d=[];l<o.length;l++)a=o[l],Object.prototype.hasOwnProperty.call(u,a)&&u[a]&&d.push(u[a][0]),u[a]=0;for(r in i)Object.prototype.hasOwnProperty.call(i,r)&&(e[r]=i[r]);f&&f(t);while(d.length)d.shift()();return c.push.apply(c,s||[]),n()}function n(){for(var e,t=0;t<c.length;t++){for(var n=c[t],r=!0,a=1;a<n.length;a++){var o=n[a];0!==u[o]&&(r=!1)}r&&(c.splice(t--,1),e=i(i.s=n[0]))}return e}var r={},a={app:0},u={app:0},c=[];function o(e){return i.p+"js/"+({}[e]||e)+"."+{"chunk-0447b482":"393a876a","chunk-503fa6db":"34a1336d","chunk-d91c31e8":"bee67a8c"}[e]+".js"}function i(t){if(r[t])return r[t].exports;var n=r[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,i),n.l=!0,n.exports}i.e=function(e){var t=[],n={"chunk-503fa6db":1,"chunk-d91c31e8":1};a[e]?t.push(a[e]):0!==a[e]&&n[e]&&t.push(a[e]=new Promise((function(t,n){for(var r="css/"+({}[e]||e)+"."+{"chunk-0447b482":"31d6cfe0","chunk-503fa6db":"bd1d2ea2","chunk-d91c31e8":"3334a870"}[e]+".css",u=i.p+r,c=document.getElementsByTagName("link"),o=0;o<c.length;o++){var s=c[o],l=s.getAttribute("data-href")||s.getAttribute("href");if("stylesheet"===s.rel&&(l===r||l===u))return t()}var d=document.getElementsByTagName("style");for(o=0;o<d.length;o++){s=d[o],l=s.getAttribute("data-href");if(l===r||l===u)return t()}var f=document.createElement("link");f.rel="stylesheet",f.type="text/css",f.onload=t,f.onerror=function(t){var r=t&&t.target&&t.target.src||u,c=new Error("Loading CSS chunk "+e+" failed.\n("+r+")");c.code="CSS_CHUNK_LOAD_FAILED",c.request=r,delete a[e],f.parentNode.removeChild(f),n(c)},f.href=u;var b=document.getElementsByTagName("head")[0];b.appendChild(f)})).then((function(){a[e]=0})));var r=u[e];if(0!==r)if(r)t.push(r[2]);else{var c=new Promise((function(t,n){r=u[e]=[t,n]}));t.push(r[2]=c);var s,l=document.createElement("script");l.charset="utf-8",l.timeout=120,i.nc&&l.setAttribute("nonce",i.nc),l.src=o(e);var d=new Error;s=function(t){l.onerror=l.onload=null,clearTimeout(f);var n=u[e];if(0!==n){if(n){var r=t&&("load"===t.type?"missing":t.type),a=t&&t.target&&t.target.src;d.message="Loading chunk "+e+" failed.\n("+r+": "+a+")",d.name="ChunkLoadError",d.type=r,d.request=a,n[1](d)}u[e]=void 0}};var f=setTimeout((function(){s({type:"timeout",target:l})}),12e4);l.onerror=l.onload=s,document.head.appendChild(l)}return Promise.all(t)},i.m=e,i.c=r,i.d=function(e,t,n){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(i.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)i.d(n,r,function(t){return e[t]}.bind(null,r));return n},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/static/view/dist/",i.oe=function(e){throw console.error(e),e};var s=window["webpackJsonp"]=window["webpackJsonp"]||[],l=s.push.bind(s);s.push=t,s=s.slice();for(var d=0;d<s.length;d++)t(s[d]);var f=l;c.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"1fef":function(e,t,n){"use strict";n("7145")},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var r=n("7a23"),a={style:{height:"6%"},class:"label_style bor"},u=Object(r["p"])("文件上传"),c=Object(r["p"])("Token管理"),o=Object(r["p"])("后台管理");function i(e,t,n,i,s,l){var d=Object(r["R"])("el-menu-item"),f=Object(r["R"])("el-menu"),b=Object(r["R"])("el-aside"),p=Object(r["R"])("el-header"),h=Object(r["R"])("router-view"),m=Object(r["R"])("el-main"),g=Object(r["R"])("Foo"),v=Object(r["R"])("el-footer"),j=Object(r["R"])("el-container");return Object(r["I"])(),Object(r["k"])(j,{style:Object(r["z"])(s.container_style)},{default:Object(r["gb"])((function(){return[Object(r["q"])(b,{style:Object(r["z"])(s.aside_width)},{default:Object(r["gb"])((function(){return[Object(r["n"])("div",a,Object(r["V"])(s.label),1),Object(r["q"])(f,{router:"router",class:"el-menu-demo","background-color":"#545c64","text-color":"#fff","active-text-color":"#ffd04b"},{default:Object(r["gb"])((function(){return[Object(r["q"])(d,{index:"upload"},{default:Object(r["gb"])((function(){return[u]})),_:1}),Object(r["q"])(d,{index:"token_manager"},{default:Object(r["gb"])((function(){return[c]})),_:1}),Object(r["q"])(d,{disabled:!s.admin,index:"work_manager"},{default:Object(r["gb"])((function(){return[o]})),_:1},8,["disabled"])]})),_:1})]})),_:1},8,["style"]),Object(r["q"])(j,{style:{padding:"0"}},{default:Object(r["gb"])((function(){return[Object(r["q"])(p,{style:Object(r["z"])(s.header_height)},null,8,["style"]),Object(r["q"])(m,null,{default:Object(r["gb"])((function(){return[Object(r["q"])(h)]})),_:1}),Object(r["q"])(v,{class:"bor"},{default:Object(r["gb"])((function(){return[Object(r["q"])(g)]})),_:1})]})),_:1})]})),_:1},8,["style"])}n("b0c0");var s={style:{width:"100%",height:"100%",padding:"10px"}},l=Object(r["p"])("提交任务");function d(e,t,n,a,u,c){var o=Object(r["R"])("el-input"),i=Object(r["R"])("el-col"),d=Object(r["R"])("el-button"),f=Object(r["R"])("el-row");return Object(r["I"])(),Object(r["m"])("div",s,[Object(r["q"])(f,null,{default:Object(r["gb"])((function(){return[Object(r["q"])(i,{span:8},{default:Object(r["gb"])((function(){return[Object(r["q"])(o,{disabled:!u.admin,modelValue:u.name,"onUpdate:modelValue":t[0]||(t[0]=function(e){return u.name=e}),placeholder:"请输入任务名称"},null,8,["disabled","modelValue"])]})),_:1}),Object(r["q"])(i,{span:8},{default:Object(r["gb"])((function(){return[Object(r["q"])(o,{disabled:!u.admin,modelValue:u.end_time,"onUpdate:modelValue":t[1]||(t[1]=function(e){return u.end_time=e}),placeholder:"请输入截至时间"},null,8,["disabled","modelValue"])]})),_:1}),Object(r["q"])(i,{span:5},{default:Object(r["gb"])((function(){return[Object(r["q"])(d,{disabled:!u.admin,type:"success",onClick:c.create_work},{default:Object(r["gb"])((function(){return[l]})),_:1},8,["disabled","onClick"])]})),_:1})]})),_:1})])}var f=n("7c15"),b=n.n(f),p={name:"Foo",data:function(){return{name:"",end_time:"",admin:!1}},created:function(){var e=this;b.a.check_token().then((function(t){e.admin=200===t.code}))},methods:{create_work:function(){b.a.create_work({name:this.name,end_time:new Date(this.end_time).getTime()/1e3})}}},h=n("6b0d"),m=n.n(h);const g=m()(p,[["render",d]]);var v=g,j={components:{Foo:v},data:function(){return{label:"作业收集系统",container_style:{margin:"0 auto",height:"100%"},header_height:{padding:0,backgroundColor:"#545c64",height:"6%"},aside_width:{backgroundColor:"#545c64",padding:0,width:"200px"},admin:!1}},methods:{},computed:{},created:function(){var e=this;this.$router.push("/upload"),window.innerWidth<600&&(this.aside_width.width="0px"),b.a.check_token().then((function(t){e.admin=200===t.code}))}};n("1fef");const O=m()(j,[["render",i]]);var k=O,w=n("7864"),_=(n("c69f"),n("3ef0")),y=n.n(_),x=function(e){e.use(w["b"],{locale:y.a})},R=(n("d3b7"),n("3ca3"),n("ddb0"),n("6c02")),q=[{path:"/upload",name:"Upload",component:function(){return n.e("chunk-d91c31e8").then(n.bind(null,"2679"))}},{path:"/token_manager",name:"token_Manager",component:function(){return n.e("chunk-503fa6db").then(n.bind(null,"5dd8"))}},{path:"/work_manager",name:"work_Manager",component:function(){return n.e("chunk-0447b482").then(n.bind(null,"66f2"))}}],C=Object(R["a"])({history:Object(R["b"])(),routes:q}),S=C,P=Object(r["j"])(k).use(S);x(P),P.mount("#app")},7145:function(e,t,n){},"7c15":function(e,t,n){var r=n("c973").default;n("96cf");var a=n("bc3a"),u=n("e144"),c=u.v4;e.exports={base:"",get_files:function(){var e=r(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,a.post(this.base+"/public/get_files/"+t);case 2:return n=e.sent,e.abrupt("return",n.data.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),get_works:function(){var e=r(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,a.post(this.base+"/public/get_works");case 2:return t=e.sent,e.abrupt("return",t.data.data);case 4:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),get_work:function(){var e=r(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,a.post(this.base+"/public/get_work/"+t);case 2:return n=e.sent,e.abrupt("return",n.data.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),upload:function(){var e=r(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,a.post(this.base+"/public/upload",t);case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),get_token:function(){var e=localStorage.getItem("token");if(null===e){var t=c();return localStorage.setItem("token",t),t}return e},handRemove:function(){var e=r(regeneratorRuntime.mark((function e(t,n){var r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,a.post(this.base+"/public/remove_file/"+t+"?token="+n);case 2:return r=e.sent,e.abrupt("return",r.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t,n){return e.apply(this,arguments)}return t}(),create_work:function(){var e=r(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,a.post(this.base+"/admin/create_work?token="+this.get_token(),t);case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),delete_work:function(){var e=r(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,a.post(this.base+"/admin/delete_work/"+t+"?token="+this.get_token());case 2:return n=e.sent,e.abrupt("return",n.data);case 4:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),check_token:function(){var e=r(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,a.post(this.base+"/check_token?token="+this.get_token());case 2:return t=e.sent,e.abrupt("return",t.data);case 4:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()}},c69f:function(e,t,n){}});
//# sourceMappingURL=app.adb2f4dc.js.map