{{template "./Layout/header.html"}}
	<!-- 头部导航栏 start : -->
	<div class="navbar">
		<div class="navbar-inner">
			<div class="container-fluid">
				<a class="btn btn-navbar" data-toggle="collapse" data-target=".top-nav.nav-collapse,.sidebar-nav.nav-collapse">
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</a>
				<a id="main-menu-toggle" class="hidden-phone open"><i class="icon-reorder"></i></a>		
				<div class="row-fluid">
				<a class="brand span2" href="index.html"><span>我的GO项目</span></a>
				</div>		
				<!-- start: Header Menu -->
				<div class="nav-no-collapse header-nav">
					<ul class="nav pull-right">
						<!-- start: User Dropdown -->
						<li class="dropdown">
							<a class="btn account dropdown-toggle" data-toggle="dropdown" href="#">
								<div class="avatar"><img src="/static/img/avatar.jpg" alt="Avatar" /></div>
								<div class="user">
									<span class="hello">欢迎登录!</span>
									<span class="name">{{$.admin.Username}}</span>
								</div>
							</a>
							<ul class="dropdown-menu">
								<li class="dropdown-menu-title"></li>
								<!--<li><a href="#"><i class="icon-user"></i> Profile</a></li>
								<li><a href="#"><i class="icon-cog"></i> Settings</a></li>
								<li><a href="#"><i class="icon-envelope"></i> Messages</a></li>-->
								<li><a href="/index/logout"><i class="icon-off"></i> 退出</a></li>
							</ul>
						</li>
						<!-- end: User Dropdown -->
					</ul>
				</div>
				<!-- end: Header Menu -->
			</div>
		</div>
	</div>
	<!-- 头部导航栏 end :  -->
	<div class="container-fluid-full">
		<div class="row-fluid">	
			<!-- start: 主要导航栏 -->
			<div id="sidebar-left" class="span2">
				导航栏搜索
				<div class="row-fluid actions">							
					<input type="text" class="search span12" placeholder="搜索栏目" />
				</div>	
				
				<div class="nav-collapse sidebar-nav">
					<ul class="nav nav-tabs nav-stacked main-menu">
						<li>
							<a class="dropmenu" href="#"><i class="icon-folder-close-alt"></i><span class="hidden-tablet"> 后台管理</span> <span class="label">3</span></a>
							<ul>
								<li><a class="submenu" href="/admin/index"><i class="icon-user"></i><span class="hidden-tablet">管理员管理</span></a></li>
								<li><a class="submenu" href="/menus/index"><i class="icon-reorder"></i><span class="hidden-tablet">后台栏目管理</span></a></li>
								<li><a class="submenu" href="/other/index"><i class="icon-star"></i><span class="hidden-tablet">Icons预览</span></a></li>
							</ul>	
						</li>
						<li><a href="/category/index"><i class="icon-edit"></i><span class="hidden-tablet">分类管理</span></a></li>
					</ul>
				</div>
			</div>
			<!-- end: 主要导航栏 -->
						
			<!-- start: 主要内容 -->
			<div id="content" class="span10">