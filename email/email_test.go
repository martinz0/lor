package email

import (
	"testing"
)

func TestSend(t *testing.T) {
	if err := Send("new_options@126.com", "GOEMAIL", "text/html", `
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
	<meta http-equiv="content-Type" content="text/html; charset=gb2312" />
	<meta name="keywords" Content="网易帮助,网易在线帮助,网易自助服务,邮箱帮助,博客帮助,相册帮助,将军令帮助,一卡通帮助,邮箱客服,网易客服,在线客服,网易助手小易">
	<meta name="description" Content="网易帮助中心(help.163.com),为网易用户提供在线的产品帮助,您还可以与网易助手小易即时对话解决使用网易产品时遇到的问题">
	<meta name="author" Content="NetEase Customer Service,Customer Experience & Development Team">
	<meta name="copyright" Content="NetEase.com,Inc." />
	<meta name="robots" content="index, follow" />
	<meta name="googlebot" content="index, follow" />
	<link rel="shortcut icon" href="http://www.163.com/favicon.ico" />
	<title>企业退信的常见问题？-163邮箱常见问题</title>
	<link href="http://cimg2.163.com/help/page_common20091106.css" rel="stylesheet" type="text/css" media="all" />
	<link href="http://help.163.com/special/00754IGS/163mail_style.css" rel="stylesheet" type="text/css" />
	<!-- 
	<script type="text/javascript" src="http://cimg2.163.com/help/V2_9/javascript/zxml.js" ></script>
	<script type="text/javascript" src="https://mima.help.163.com/freemail/javascript/ts/payBill/zxml.js" ></script>
	<script type="text/javascript" src="https://mima.help.163.com/freemail/javascript/ts/payBill/index.js" ></script>
	-->
	<script type="text/javascript" src="/special/00754JSL/freemail_paybill_index.js" ></script>
	<script type="text/javascript" src="http://help.163.com/special/00754IGS/guidepath.js" ></script>
	<script type="text/javascript" src="http://help.163.com/special/00754IGS/pagejs.js" ></script>
	</head>
	<body class="body-bg">
	<div class="wrap">
	<div class="header">
	<div class="header-logo"><a class="logo1" href="http://mail.163.com/" target="_blank"></a> <a class="logo2" href="http://help.163.com/" target="_blank"></a> </div>
	<div class="header-link">
	<ul>
	<li><a href="http://mail.163.com/" target="_blank">&nbsp;&nbsp;邮箱首页</a></li>
	<li class="hd-width"><a href="http://vip.163.com/" target="_blank">VIP163尊贵邮</a></li>
	<li class="hd-width"><a href="http://vip.126.com/" target="_blank">VIP126尊享邮</a></li>
	<li><a href="http://www.188.com/" target="_blank">188财富邮</a></li>
	<li><a href="http://mail.126.com/" target="_blank">126免费邮</a></li>
	<li class="hd-b"><a href="http://www.yeah.net/" target="_blank">Yeah.net</a></li>
	</ul>
	</div>
	</div>
	<!--/header end-->
	<div class="con">
	<div class="nav">
	<ul id="topTabs">
	<li><a href="/special/007525G0/163mail_index.html">首&nbsp;&nbsp;&nbsp;页</a></li>
	<li><a href="/special/007525G0/163mail_faq.html">问题分类</a></li>
	<li><a href="/special/007525G0/163mail_ans.html">在线提问</a></li>
	<li><a href="/special/007525G0/zzcxlogin_163.html">自助查询</a></li>
	<li><a href="http://help.mail.163.com/" target="_blank">邮箱知道</a></li>
	<li><a href="http://email2.163.com/info.htm" target="_blank">网络检测</a></li>
	</ul>
	</div>
	<!--/nav end-->
	<script type="text/javascript">
	topTabOn(2);
	var guideUrl='http://help.163.com/special/007525G0/163mail_guide.html';
	</script>
	<div class="nav_bottom">
	<div class="nb_text"> 您当前所处的位置： <a href="/special/007525G0/163mail_index.html">帮助中心</a> <span id="guidePathBox"></span></div>
	</div>
	<!--/nav_bottom end-->
	<div class="clear"></div>
	<!--中部主体 开始-->
	<div class="content">
	<div class="Fl_left">
	<h2>请选择：</h2>
	<ul id="lftTreeUl">
	</ul>
	<!--目录树下方广告-->
	<!--广告1 收费免费对比推广AD-->
	<div id="AD_1" class="ad">     <a href="http://vip.163.com/vip/special/compare.htm?from=163help_im" target="_blank">
	<img src="http://cimg2.163.com/help/vip/img/vipvsfree03.jpg"  border="0" title="升级VIP邮箱，尊享20多项特权" />
	</a>
	</div>
	<!--广告2 电子传真推广AD-->
	<div id="AD_2" class="ad">     <a href="http://fax.163.com?163help_im" target="_blank">
	<img src="http://cimg2.163.com/help/urs/img/faxad12112001.png"  border="0" title="免传真机，上网即可收发传真" />
	</a>
	</div>
	<!--广告3 随身邮推广AD-->
	<div id="AD_3" class="ad" style="display:none;">
	<a href="http://jf.vip.163.com/apps/upgrade/unauth/index.m?163help/sms" target="_blank">
	<img src="http://cimg2.163.com/help/163mail/img/163ssy.jpg"  border="0"/>
	</a>
	</div>
	<!--广告4 升级VIP推广AD-->
	<div id="AD_4" class="ad" style="display:none;">
	<a href="http://jf.vip.163.com/apps/upgrade/unauth/index.m?163help/entry" target="_blank">
	<img src="http://cimg2.163.com/help/163mail/img/163dljs.jpg"  border="0"/>
	</a>
	</div>
	<!--广告5 文件恢复推广AD-->
	<div id="AD_5" class="ad" style="display:none;">
	<a href="http://jf.vip.163.com/apps/upgrade/unauth/index.m?163help/lossmail" target="_blank">
	<img src="http://cimg2.163.com/help/163mail/img/163yjds.jpg"  border="0"/>
	</a>
	</div>
	<!--广告6 升级附件推广AD-->
	<div id="AD_6" class="ad" style="display:none;">
	<a href="http://jf.vip.163.com/apps/upgrade/unauth/index.m?163help/appendix" target="_blank">
	<img src="http://cimg2.163.com/help/163mail/img/163kj.jpg"  border="0"/>
	</a>
	</div>
	<!--广告7 登录保障推广AD-->
	<div id="AD_7" class="ad" style="display:none;">
	<a href="http://jf.vip.163.com/apps/upgrade/unauth/index.m?163help/safe" target="_blank">
	<img src="http://cimg2.163.com/help/163mail/img/163ybz.jpg"  border="0" />
	</a>
	</div>
	<script type="text/javascript">
	(function(win){
		var _win = win,
		_doc = _win.document,
		_$ = function(id){ return _doc.getElementById(id); },
		_getArgs = _win['GetArgsFromHref'];

		var AD_config = {
			2805 : ['AD_3'],
			2471 : ['AD_4'],
			2558 : ['AD_5'],
			3252 : ['AD_6'],
			2182 : ['AD_7']
		};

		function changeAD() {
			var now = new Date();
			_$('ad_onlineOrder').style.display = 'none';
			_$('ad_onlineUpgrade').style.display = 'block';
			// 周一到周五 10：00-12:00   14:00-17:00
			if(now.getDay() != 0 && now.getDay() != 6) {
				if( (now.getHours() >= 10 && now.getHours() <12) 
				|| (now.getHours() >= 14 && now.getHours() <17) ) {
					_$('ad_onlineUpgrade').style.display = 'none';
					_$('ad_onlineOrder').style.display = 'block';
				}
			}
		}
		function showAD() {
			var xid = _getArgs('id'),
			ad_num = 0;
			if( xid == '' ) {
				var pathLength = _$('guidePathBox').getElementsByTagName('a').length,
				lastPath = _$('guidePathBox').getElementsByTagName('a')[pathLength-1];
				if( lastPath !== undefined ) xid = lastPath.getAttribute('id').replace(/guide/g, '');
			}
			if ( xid != '' ) {
				for(var i in AD_config[xid]) {
					if ( _$(AD_config[xid][i]) !== undefined  ) {
						_$(AD_config[xid][i]).style.display="block";
						ad_num++;
					}
				}
				if(ad_num > 0) {
					_$('AD_1').style.display="none";
					_$('AD_2').style.display="none";
				}
			}
		}
		_win['AD'] = {};
		_win['AD']['showAD'] = showAD;
		_win['AD']['changeAD'] = changeAD;
	})(window);
	elementAttchEvent(window,'load',AD.showAD);
	// elementAttchEvent(window,'load',AD.changeAD);
	</script>
	</div>
	<div class="Fl_right">
	<div class="searchArea">
	<form onsubmit="return search_proxy(this);" target="_blank">
	<input id="searchKey" type="text" value="输入关键词（如：修复密码、一键升级）点击搜索" class="schInp" onfocus="inEvt('focus',this)" onblur="inEvt('blur',this)" />
	<input type="submit" value=" 搜 索 " class="submit_sch" />
	</form>
	<div class="keyword">
	<span>热门咨询：
	&nbsp;&nbsp;</span><a style="color:#cc0000" href="http://help.163.com/10/1108/18/6L03I4G500753VB8.html?163help_search" target="_blank">iPhone设置</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://help.163.com/09/1215/11/5QIOSVIO00753VB8.html?sousuo" target="_blank">密码修复</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://help.163.com/09/1216/17/5QLVS5CT00753VB8.html?sousuo" target="_blank" style="color:#000">无法登录</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://help.163.com/09/1218/10/5QQFVP1F00753VB8.html?sousuo" target="_blank" style="color:#000">邮件丢失</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://help.163.com/12/0814/13/88SFNJTR00753VB8.html?sousuo" target="_blank" style="color:#000">异地登录</a>
	<!--&nbsp;&nbsp;</span><a href="http://reg.vip.163.com/upgradeIndex.m?163help_gjz" target="_blank">升级为vip</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://vip.163.com/vip/special/vipts/vipts.htm?4,9#?163help_gjz" target="_blank">邮箱管家</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://vip.163.com/vip/special/vipts/vipts.htm?6,19#?163help_gjz" target="_blank" style="color:#000">电子传真</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://vip.163.com/vip/special/vipts/vipts.htm?2,2#?163help_gjz" target="_blank" style="color:#000">400群发</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="http://help.163.com/09/1217/10/5QNSANC500753VB8.html?sousuo" target="_blank" style="color:#000">收不到邮件</a>-->
	</div>
	<script type="text/javascript">
	var serachUrl="http://help.163.com/special/007525G0/urs-163.html?s=1#start=0&q=";				
	function doSearch(oForm){
		var key=document.getElementById('searchKey').value;
		key=key.replace(/(^\s*)|(\s*$)/g, "");
		if(key=="输入关键字 如：163邮箱 登录" || key=="" ){
			alert("请输入关键字,如：163邮箱 登录");
		}else{
			var url=serachUrl+encodeURIComponent(" "+key);
			if(document.all){
				var objA=document.createElement("A");
				objA.href=url;
				objA.target="_blank";
				oForm.appendChild(objA);
				objA.click();
				oForm.removeChild(objA);
			}else{
				window.open(url);
			}
			//oForm.removeChild(objA);
			//return true;	
		}
		return false;
	}
	function search_proxy(oForm) {
		var keyword=document.getElementById('searchKey').value;
		keyword=keyword.replace(/(^\s*)|(\s*$)/g, "");
		if ( /^6\d{6}$/.test(keyword) ) {
			return search_servCode(keyword, oForm, doSearch);
		}else if ( /^\d{6}$/.test(keyword) ) {
			return search_old580(keyword, oForm, doSearch);
		}else{
			return doSearch(oForm);
		}

	}
	</script>
	</div>
	<!--/searchArea end-->
	<div class="page_content">
	<h2>企业退信的常见问题？</h2>
	<div class="btn_fh">
	<input type="button" onclick="window.history.go(-1)" value="返回" class="fanhui"/>
	</div>
	<div class="con_text"> 
	<p>网易鼓励企业用户在遵守互联网标准的前提下，向网易邮箱用户发送已经许可的邮件。为了确保您的邮件能及时准确的投递到网易邮箱，请首先确认您没有违背网易的反垃圾政策，点击<a href="http://help.163.com/09/1224/14/5RAB4VK500753VB8.html" target="_blank">查看详情</a>。<br  />　　若您的邮件仍无法到达网易邮箱，并收到退信，请根据退信的返回字段，在下面的表单查询具体的退信原因。如果这些信息仍不能解决你的问题，那么请<a href="http://feedback.mail.163.com/" target="_blank">点击这里填写故障反馈表</a>。<br  />
	<strong>退信代码说明：</strong>
	<br  />　　&#8226;421 HL:REP 该IP发送行为异常，存在接收者大量不存在情况，被临时禁止连接。请检查是否有用户发送病毒或者垃圾邮件，并核对发送列表有效性；<br  />　　&#8226;421 HL:ICC 该IP同时并发连接数过大，超过了网易的限制，被临时禁止连接。请检查是否有用户发送病毒或者垃圾邮件，并降低IP并发连接数量；<br  />　　&#8226;421 HL:IFC 该IP短期内发送了大量信件，超过了网易的限制，被临时禁止连接。请检查是否有用户发送病毒或者垃圾邮件，并降低发送频率；<br  />　　&#8226;421 HL:MEP 该IP发送行为异常，存在大量伪造发送域域名行为，被临时禁止连接。请检查是否有用户发送病毒或者垃圾邮件，并使用真实有效的域名发送；<br  />　　&#8226;450 MI:CEL 发送方出现过多的错误指令。请检查发信程序；<br  />　　&#8226;450 MI:DMC 当前连接发送的邮件数量超出限制。请减少每次连接中投递的邮件数量；<br  />　　&#8226;450 MI:CCL 发送方发送超出正常的指令数量。请检查发信程序；<br  />　　&#8226;450 RP:DRC 当前连接发送的收件人数量超出限制。请控制每次连接投递的邮件数量；<br  />　　&#8226;450 RP:CCL 发送方发送超出正常的指令数量。请检查发信程序；<br  />　　&#8226;450 DT:RBL 发信IP位于一个或多个RBL里。请参考<a href="http://www.rbls.org/">http://www.rbls.org/</a>关于RBL的相关信息；<br  />　　&#8226;450 WM:BLI 该IP不在网易允许的发送地址列表里；<br  />　　&#8226;450 WM:BLU 此用户不在网易允许的发信用户列表里；<br  />　　&#8226;451 DT:SPM ,please try again 邮件正文带有垃圾邮件特征或发送环境缺乏规范性，被临时拒收。请保持邮件队列，两分钟后重投邮件。需调整邮件内容或优化发送环境；<br  />　　&#8226;451 Requested mail action not taken: too much fail authentication 登录失败次数过多，被临时禁止登录。请检查密码与帐号验证设置；<br  />　　&#8226;451 RP:CEL 发送方出现过多的错误指令。请检查发信程序；<br  />　　&#8226;451 MI:DMC 当前连接发送的邮件数量超出限制。请控制每次连接中投递的邮件数量；<br  />　　&#8226;451 MI:SFQ 发信人在15分钟内的发信数量超过限制，请控制发信频率；<br  />　　&#8226;451 RP:QRC 发信方短期内累计的收件人数量超过限制，该发件人被临时禁止发信。请降低该用户发信频率；<br  />　　&#8226;451 Requested action aborted: local error in processing 系统暂时出现故障，请稍后再次尝试发送；<br  />　　&#8226;500 Error: bad syntaxU 发送的smtp命令语法有误；<br  />　　&#8226;550 MI:NHD HELO命令不允许为空；<br  />　　&#8226;550 MI:IMF 发信人电子邮件地址不合规范。请参考<a href="http://www.rfc-editor.org/">http://www.rfc-editor.org/</a>关于电子邮件规范的定义；<br  />　　&#8226;550 MI:SPF 发信IP未被发送域的SPF许可。请参考<a href="http://www.openspf.org/">http://www.openspf.org/</a>关于SPF规范的定义；<br  />　　&#8226;550 MI:DMA 该邮件未被发信域的DMARC许可。请参考<a href="http://dmarc.org/">http://dmarc.org/</a>关于DMARC规范的定义；<br  />　　&#8226;550 MI:STC 发件人当天的连接数量超出了限定数量，当天不再接受该发件人的邮件。请控制连接次数；<br  />　　&#8226;550 RP:FRL 网易邮箱不开放匿名转发（Open relay）；<br  />　　&#8226;550 RP:RCL 群发收件人数量超过了限额，请减少每封邮件的收件人数量；<br  />　　&#8226;550 RP:TRC 发件人当天内累计的收件人数量超过限制，当天不再接受该发件人的邮件。请降低该用户发信频率；<br  />　　&#8226;550 DT:SPM 邮件正文带有很多垃圾邮件特征或发送环境缺乏规范性。需调整邮件内容或优化发送环境；<br  />　　&#8226;550 Invalid User 请求的用户不存在；<br  />　　&#8226;550 User in blacklist 该用户不被允许给网易用户发信；<br  />　　&#8226;550 User suspended 请求的用户处于禁用或者冻结状态；<br  />　　&#8226;550 Requested mail action not taken: too much recipient&nbsp; 群发数量超过了限额；<br  />　　&#8226;552 Illegal Attachment 不允许发送该类型的附件，包括以.uu .pif .scr .mim .hqx .bhx .cmd .vbs .bat .com .vbe .vb .js .wsh等结尾的附件；<br  />　　&#8226;552 Requested mail action aborted: exceeded mailsize limit 发送的信件大小超过了网易邮箱允许接收的最大限制；<br  />　　&#8226;553 Requested action not taken: NULL sender is not allowed 不允许发件人为空，请使用真实发件人发送；<br  />　　&#8226;553 Requested action not taken: Local user only&nbsp; SMTP类型的机器只允许发信人是本站用户；<br  />　　&#8226;553 Requested action not taken: no smtp MX only&nbsp; MX类型的机器不允许发信人是本站用户；<br  />　　&#8226;553 authentication is required&nbsp; SMTP需要身份验证，请检查客户端设置；<br  />　　&#8226;554 DT:SPM 发送的邮件内容包含了未被许可的信息，或被系统识别为垃圾邮件。请检查是否有用户发送病毒或者垃圾邮件；<br  />　　&#8226;554 DT:SUM 信封发件人和信头发件人不匹配；<br  />　　&#8226;554 IP is rejected, smtp auth error limit exceed 该IP验证失败次数过多，被临时禁止连接。请检查验证信息设置；<br  />　　&#8226;554 HL:IHU 发信IP因发送垃圾邮件或存在异常的连接行为，被暂时挂起。请检测发信IP在历史上的发信情况和发信程序是否存在异常；<br  />　　&#8226;554 HL:IPB 该IP不在网易允许的发送地址列表里；<br  />　　&#8226;554 MI:STC 发件人当天内累计邮件数量超过限制，当天不再接受该发件人的投信。请降低发信频率；<br  />　　&#8226;554 MI:SPB 此用户不在网易允许的发信用户列表里；<br  />　　&#8226;554 IP in blacklist 该IP不在网易允许的发送地址列表里。</p>
	<div id="pageGGBox"></div>
	</div>
	</div>
	<div id="suc" class="survey">
	<h5>您对以上帮助信息感到：</h5>
	<ul>
	<li id="s0"><a href="javascript:show_comment(0);" onfocus="this.blur()">非常满意</a></li>
	<li id="s1"><a href="javascript:show_comment(1);" onfocus="this.blur()">满意</a></li>
	<li id="s2"><a href="javascript:show_comment(2);" onfocus="this.blur()">一般</a></li>
	<li id="s3"><a href="javascript:show_comment(3);" onfocus="this.blur()">不满意</a></li>
	<li id="s4"><a href="javascript:show_comment(4);" onfocus="this.blur()">非常不满意</a></li>
	</ul>
	<div class="survey_con" id="sur_con" style="display:none;">
	<div class="close" id="sur_tips">
	<!-- a href="javascript:hide_close();"  onfocus="this.blur()"><b></b>关闭</a -->
	</div>
	<div class="survey_con_suc" id="sur0" >
	<p>感谢您的反馈，我们已记录，有您的支持我们会做的更好！谢谢！</p>
	<br>
	</div>
	<div id="sur1">
	<p id="two_TX" onfocus="textCl('focus',this)" onblur="textCl('blur',this)" ></p>

	<a class="link_tg" href="javascript:show_comment(1);" onfocus="this.blur()" id="skip" style="display:none;"> </a>

	</div>
	</div>
	</div>
	<script>
	var obj = function(id) { return document.getElementById(id); }
	String.prototype.trim = function() { return this.replace(/(^\s*)|(\s*$)/g, ""); }
	function SHOW(id) { obj(id).style.display = "block"; }
	function HIDE(id) { obj(id).style.display = "none"; }
	var hf='';
	var reson='';
	var prj="faq";
	var surLink="http://guide.help.163.com/xiaoyi/survey.php?issubmit=3&pj="+prj+"&u="+encodeURI(window.location.href);
	function surAddInfo(){
		var surUrl=surLink+"&c="+encodeURI(hf)+"&r="+encodeURI(reson);
		var sg=obj('two_TX').value.trim();
		if(sg!='' && sg!=two_tx[0] && sg!=two_tx[1]){
			surUrl+="&s="+encodeURI(sg);
		}
		//alert(surUrl);
		var iMg=new Image();
		iMg.onload=function(){};
		iMg.src=surUrl;
	}
	var suc_flag=false;
	//var two_tx=["请描述您不明白的地方或提出您的建议？","请描述您所遇到的问题或提出您的建议？"];
	function hide_close(){
		obj("sur_con").style.display="none";
		for(var i=0; i<=4;i++){
			obj("s"+i).className="";
		}
	}
	function show_con(m){
		var sksksk=["有帮助","我看不明白","不是我要的答案"];
		var a=isSubmitSe();
		if(a){
			if(m==0){
				suc_flag=true;
				hf=1;
				reson='';
				surAddInfo();
			}else{
				hf=0;
			}
			reson=sksksk[m];
			for(var i=0; i<3;i++){
				obj("s"+i).className="suroff";
				if(i<=1){obj("sur"+i).style.display="none";}
			}

			obj("s"+m).className="suron";
			if(m>0){
				var TX=m-1;
				obj("two_TX").value=two_tx[TX];
			}
			m=m>1?1:m;
			obj("sur"+m).style.display="block";
			obj("sur_con").style.display="block";
		}
	}
	function show_suc(){
		for(var i=0; i<2;i++){obj("sur"+i).style.display="none";}
		obj("sur"+0).style.display="block";
		suc_flag=true;surAddInfo();
	}
	function isSubmitSe(){
		if(suc_flag==true){
			alert("您的信息已提交");
			return false;
		}
		return true;
	}
	//new survey
	var level = ["非常满意", "满意", "一般", "不满意", "非常不满意"];
	var tipss = ["谢谢您！请问是否要添加其他内容？", "很遗憾没帮助到您，让客服来处理您的问题吧！&nbsp;&nbsp;&nbsp;&nbsp;<a href=\"http://help.mail.163.com/feedback.do?fromhelp\" target=\"_blank\">在线反馈>></a>"];
	var two_tx = ["我们会根据您描述的问题仔细核实，对帮助进行优化，感谢您的反馈！", 
	"也可发在反馈，我们会根据您描述的问题仔细核实，对帮助进行优化，感谢您的反馈！"];
	function surveyOn(m){
		for(var i=0; i<=4;i++){
			obj("s"+i).className="suroff";
		}
		obj("s"+m).className="suron";
	}
	function setHf(m) {
		if(m<=2) {
			hf = 1;
		}else{
			hf = 0;
		}
	}
	//是否显示填写框
	function show_textarea(m) {
		if(m>=2) { 
			SHOW('two_TX');
			obj('sur1').style.height = '140px';
		}else{
			HIDE('two_TX');
			obj('sur1').style.height = '40px';
		}
	}
	function show_tips(m) {
		if(m <= 1) { 
			//obj("two_TX").value = two_tx[0];
			obj("sur_tips").innerHTML = "";
		}else if(m == 2){
			obj("sur_tips").innerHTML = "<p1>" + tipss[0] + "</p1>";
			obj("two_TX").value = two_tx[0];
		}else{
			obj("sur_tips").innerHTML = "<p1>" + tipss[1] + "</p1>";
			obj("two_TX").value = two_tx[1];
		}
	}
	function comment_required() {
		var temp = obj("two_TX").value.trim();
		if(reson==level[0] || reson==level[1] || reson==level[2]) {
			return false;
		}else{
			if(temp == two_tx[0] || temp == two_tx[1] || temp == '') {
				alert("请填写您不满意原因");
				return true;
			}else{
				return false;
			}
		}
	}
	function show_comment(m) {
		if(isSubmitSe()) {
			if(m==2){
				surveyOn(m);
				SHOW("sur_con");
				HIDE("sur0");
				SHOW("sur1");
				//SHOW("skip");
				show_tips(m);
				reson=level[m];
				setHf(m);     
			}else{
				surveyOn(m);
				SHOW("sur_con");
				HIDE("sur0");
				SHOW("sur1");
				//HIDE("skip");
				//show tips
				show_tips(m);
				//show textarea
				//show_textarea(m);
				//set reson value
				reson=level[m];
				//set hf value
				setHf(m);
			}
			if(m==0 || m== 1){
				show_success();
			}
		}
	}
	function show_success(){
		if( !comment_required() ) {
			var temp = obj("two_TX").value.trim();
			if (reson == '一般' && ( temp == two_tx[0] || temp == '') ) {
				show_comment(1);
			} else {
				//clear tips
				obj("sur_tips").innerHTML = '';
				SHOW("sur_con");
				SHOW("sur0");
				HIDE("sur1");
				//for(var i=0; i<=4; i++) { obj("s"+i).className="suroff"; }
				suc_flag=true;
				//alert(reson);
				surAddInfo();
			}
		}
	}
	//show_comment(0);

	</script>
	<div class="abouttext">
	<h4><img src="http://cimg2.163.com/help/vip/img/jt_icon.jpg" />&nbsp;&nbsp;其他相关信息</h4>
	<ul id="relatePagesBox">
	<li>目前没有相关文章</li>
	</ul>
	</div>
	</div>
	</div>
	<div class="clear"></div>
	<div class="clear"></div>
	</div><!--/con end-->
	<div class="Footer"><a href="http://gb.corp.163.com/gb/home.shtml" target="_blank">关于网易</a>　<a href="http://help.163.com/?163help" target="_blank">客户服务</a>　<a href="http://mail.blog.163.com/" target="_blank">邮箱官方博客</a>  <a href="http://vipmail.163.com/" target="_blank">网易VIP邮箱</a>　<a href="http://cards.163.com/" target="_blank">精美贺卡</a>　<a href="http://help.163.com/12/0911/17/8B4TSKV400754K4E.html" target="_blank">网易电子传真</a>　<a href="http://corp.163.com/gb/legal/legal.html" target="_blank">隐私政策</a>　<br />　网易公司版权所有 &copy; 1997-<script src="http://mimg.126.net/copyright/year.js" language="javascript"></script></div>	
	</div><!--/#wrap-->
	<script type="text/javascript" src="http://analytics.163.com/ntes.js"></script>
	<script type="text/javascript">
	_ntes_nacc = "help";
	neteaseTracker();
	</script>
	<script type="text/javascript" src="http://mimg.127.net/p/tools/jquery/jquery-1.7.min.js"></script>
	<script type="text/javascript" src="/special/00754IL6/articleinlinetitle.js"></script>
	<script type="text/javascript" src="/special/00754IL6/freemailsearch.js"></script>
	<script type="text/javascript" src="/special/00754IL6/freemailgoback.js"></script>
	<!-- START searchServCode -->
	<script type="text/javascript" src="/special/00752FD5/srvcode2.js"></script>
	<script type="text/javascript" src="/special/00754IO5/servcode_urs.js"></script>
	<script type="text/javascript" src="/special/00754IGS/servcode_163.js"></script>
	<script type="text/javascript" src="/special/00754IGR/servcode_126.js"></script>
	<script type="text/javascript" src="/special/00754IGQ/servcode_yeah.js"></script>
	<script type="text/javascript" src="/special/00754IV8/servcode_vip163.js"></script>
	<script type="text/javascript" src="/special/00754IJB/servcode_vip126.js"></script>
	<script type="text/javascript" src="/special/00754IN5/servcode_188.js"></script>
	<script type="text/javascript" src="/special/00754K4G/servcode_fax.js"></script>
	<script type="text/javascript" src="/special/007541LP/searchservcode.js"></script>
	<script type="text/javascript" src="/special/007541LP/addservcode2link.js"></script>
	<!-- END searchServCode -->
	</body>
	</html>
	<!--飘窗广告-->
	<div id="pc" class="pc">
	<div><a class="pc_gb" title="关闭" onclick="document.getElementById('pc').style.display='none'"></a></div>
	<div class="pc_bt">
	<!--<a href="http://fax.163.com?163help_pc" target="_blank"><span class="pc_btico"></span></a> --> 
	<a href="http://126.am/yixinhelp" target="_blank"><span class="pc_btico"></span></a>
	</div>
	</div>
	<script type="text/javascript">
	var now = new Date();
	var day = now.getDay();
	var hour = now.getHours();
	var minute = now.getMinutes();
	if(day>=1 && day<=5 && ((hour>=10 && hour<17) || (hour==9 && minute>=30) || (hour==17 && minute<=30))){
		document.getElementById("pc").style.display="block";	
	}
	else
	document.getElementById("pc").style.display="block";	 
	</script>
	<!--飘窗广告 end-->
	<!--小易飘窗-->
	<div id="pc_1" class="pc_1">	
	<a href="http://help.163.com/special/007525G0/163mail_ans.html?163help_pc" target="_blank"><span class="pc_1"></span></a>  
	</div>
	<script type="text/javascript">
	var now = new Date();
	var day = now.getDay();
	var hour = now.getHours();
	var minute = now.getMinutes();
	if(day>=1 && day<=7 && ((hour>=9 && hour<20) || (hour==9 && minute>=00) || (hour==20 && minute<=00))){
		document.getElementById("pc_1").style.display="block";	
	}
	else
	document.getElementById("pc_1").style.display="none";	 
	</script>
	<!--小易飘窗 end-->
	`); err != nil {
		t.Fatal(err)
	}
}
