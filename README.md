# 红岩网校Web寒假考核-豆瓣

接口文档 https://documenter.getpostman.com/view/18429077/UVkiRxpg

https部署地址 https://www.poordouban.work （进去如果资源没有加载需要F12信任并继续访问资源，详细原因下面有说）

http部署地址 http://42.192.155.29/

后端项目地址1 https://github.com/L2ncE/Douban（这个是在建组织之前的仓库，大部分的提交都在上面，后期没有再更新）

总项目地址(后端项目2，前端项目1）https://github.com/RedRock-Vacation-Assessment-Douban

## 前言💞

和前端同学一起花了不少时间完成的这个项目，虽然还是有许多不足的地方，但还是有我们自己的努力在里面。文档中有一些功能的说明，希望学长们可以看看。

## 实现功能列表🤣

*粗体为基础功能要求*

1. **账号注册登录，保持登录状态，退出登录**
2. **主页搜索功能**
3. **主页推荐功能**
4. **电影页展示评分以及统计**
5. **电影页想看/看过功能(点击之后个人页面也会更新)**
6. **电影页剧照，剧情简介，演职员（可以点击进入影人页面）**
7. **影评和短评功能**
8. **讨论区**
9. **影人页面(参考的官网)(可以点击进入作品页面)**
10. **排行榜页面**
11. **分类找电影,分了类型和国家两类**
12. **个人页面 自我介绍 想看/看过 我的影评**
13. 更改密码
14. JWT登录
15. 对入参进行检验，如用户名长度，密码长度
16. 匿名进行评论功能
17. 部署在了云服务器上，项目可以通过url正常访问并使用
18. 使用了https加密（不算完美）
19. 考虑了部分安全性 ~~不会出现教务在线更改学号可以直接开盒的状况~~
20. 密保找回密码功能
21. 移动端适配

## 接口已实现未使用列表😃

1.进入讨论区详细页

![image-20220219134336649](https://s2.loli.net/2022/02/19/xNDvjIb9KXwWCu8.png)



2.进入影评详细页

![image-20220219134430719](https://s2.loli.net/2022/02/19/iquWVjS6QgIwaFx.png)

![image-20220219134441749](https://s2.loli.net/2022/02/19/teoFvyaNkYbIpKc.png)

3.分类排行榜（实现逻辑下面有说明）

![image-20220219134526386](https://s2.loli.net/2022/02/19/TpruGK9Fc3iqVPH.png)

## 真真假假🤨

因为数据太多电影影人之类的太多，没有办法实现所有的数据都完善，所以不可避免的会有假数据出现。

### 电影的真假

真电影可以点进去，假电影因为没有它的数据所以就不能点进它的详细页，一共有三十部所有数据（包括影评，短评，讨论区等）完善的真电影。

**轮播图一正在热映:1-35 其中1-10为真电影**，也就是第一页和第二页轮播图

![image-20220217121208960](https://s2.loli.net/2022/02/17/fJBeWa7G6MZ5yts.png)

**轮播图二最近热门电影:36-85 其中36-40 61-65为真电影**，也就是第一页轮播图

![image-20220217121400902](https://s2.loli.net/2022/02/17/wTrj3MIGWoxaDh5.png)

**轮播图三最近热门电视剧:86-135 其中86-90 111-115为真电影**，也就是第一页轮播图

![image-20220217121433334](https://s2.loli.net/2022/02/17/jQ1h5Vd6o9gkJ7A.png)

### 影人的真假

每一部真电影里都会有一个真影人，因为影人需要传的数据，URL之类的更多，所以还是只有30个。点击电影详细页的影人都会跳转到那一个真影人，到影人详细页点击电影则会跳转回真电影。

## 一些功能的实现😋

很多都是前后端沟通的时候写的,很杂,也不太正式

### 分类页面与排行榜页面的分类排行榜

#### 关于分类页面

![](https://s2.loli.net/2022/02/08/IvD8c3Fnejyas9b.png)

只了做这两部分

要分几种情况

接口1./classify

接口2./classify/:type/:country 两个都选,两个都传

接口3./classify1/:country (Param) 选国家

接口4./classify2/:type (Param) 选类型

1.进入分类是全部类型全部地区,使用直接传出所有真电影的接口1 /classify

2.类型和国家都进行选择分类 使用接口2 比如选择剧情,美国 就是/classify/剧情/美国

3.类型选择全部类型 地区进行其他选择 使用接口3 比如选择美国 就是/classify1/美国

4.国家选择全部国家 类型进行其他选择 使用接口4 比如选择剧情 就是/classify2/剧情

因为只有真电影有类型国家这些,所以这部分只有真的,会传 Id 名字(Name) URL 分数(Score) 四个部分

#### 关于分类排行榜

也就是排行榜的这部分

![](https://s2.loli.net/2022/02/08/Cxik8TLnOGZ4slE.png)

他点进去是这样

![](https://s2.loli.net/2022/02/08/snAMc8RlO2jfXUy.png)

我们只要一部分

差不多是下面这样

![](https://s2.loli.net/2022/02/08/K2yhqTMcQdpFCmV.png)

接口是/classifyrank/:type   (直接后面是中文的类型名称)

最后会以评分排名(以Score从高到低排)给出以下数据(以下为例子)

![image-20220208125057408](https://s2.loli.net/2022/02/08/Upz8Ng91k6fMdXs.png)

基本和官网一样

### https与http

我们的项目有https加密和普通http两种地址，原因是我搞了半天也没找到免费能用的IP地址SSL证书，我也不知道有什么其他办法能把我的后端项目部署在https下，因此只能采取这样的办法

#### http

http://42.192.155.29/这个是我的云服务器IP地址，后端项目部署在http://42.192.155.29:8080/

#### https

https://www.poordouban.work 此网址部署了SSL证书，但是后端项目依然是http的资源，所以无法显示。因为IP地址的SSL不好搞，所以只能使用一个非正规的IP地址证书，直接访问会显示不安全并且不能访问资源

![image-20220217140253786](https://s2.loli.net/2022/02/17/TNBogU634n9C7cJ.png)

需要信任资源，继续访问  在F12界面点击后面超链接进行信任

![image-20220217140535197](https://s2.loli.net/2022/02/17/8deREa35wYGHoxs.png)

![image-20220217140503233](https://s2.loli.net/2022/02/17/uK9RweDMBGNbrXj.png)

继续访问后就没有问题了，~~虽然这个https不太完美，但也算是https吧~~

### 个人页面

个人页面可以看到自己的影评和想看、看过部分，还有自我介绍

#### 影评与想看、看过

在电影页面撰写影评并发布之后，个人页面会自动更新，个人页面的影评板块还有删除功能，可以删除自己发布过的影评

![image-20220219133214213](https://s2.loli.net/2022/02/19/XZucvlsNw6k2fPi.png)

#### 个人介绍（个人留言板）

通过发布新留言更改自己的个人介绍（个人留言板）

![image-20220219133417363](https://s2.loli.net/2022/02/19/pZw4oQPuFH3IbAl.png)

留言后效果

![image-20220219133442976](https://s2.loli.net/2022/02/19/dPktI3vYpB9zixZ.png)

## 电影ID NAME URL对照表🤗

**轮播图一正在热映:1-35 其中1-10为真电影**

**轮播图二最近热门电影:36-85 其中36-40 61-65为真电影**

**轮播图三最近热门电视剧:86-135 其中86-90 111-115为真电影**

| ID   | Name                                                         | URL                                            |
| :--- | :----------------------------------------------------------- | ---------------------------------------------- |
| 1    | 黑客帝国：矩阵重启 The Matrix Resurrections                  | https://i.postimg.cc/rwMvT2Vt/1.webp           |
| 2    | 爱情神话                                                     | https://i.postimg.cc/65r9rH9d/p2772925591.webp |
| 3    | 东北虎                                                       | https://i.postimg.cc/2jG0zhD7/p2812275146.webp |
| 4    | 雄狮少年                                                     | https://i.postimg.cc/JhjQSBxJ/p2702755317.webp |
| 5    | 魔法满屋                                                     | https://i.postimg.cc/SsqG4N1P/p2807936075.webp |
| 6    | 穿过寒冬拥抱你                                               | https://i.postimg.cc/CxqCffWt/p2801718909.webp |
| 7    | 误杀2                                                        | https://i.postimg.cc/50xk33Jj/p2770857575.webp |
| 8    | 李茂扮太子                                                   | https://i.postimg.cc/zv4j9wsC/p2812626447.webp |
| 9    | 反贪风暴5：最终章 G風暴                                      | https://i.postimg.cc/hjj1v3g8/p2686021044.webp |
| 10   | 以年为单位的恋爱                                             | https://i.postimg.cc/Y2Qz9gBs/p2715903096.webp |
| 11   | 汪汪队立大功大电影 Paw Patrol: The Movie                     | https://i.postimg.cc/ZYLv2LS4/p2808318200.webp |
| 12   | 最初的梦想                                                   | https://i.postimg.cc/fTTSN0zX/p2766852789.webp |
| 13   | 芭比公主历险记 Barbie Princess Adventure                     | https://i.postimg.cc/rmNz6q7Q/p2852973691.webp |
| 14   | 四海                                                         | https://i.postimg.cc/0QkPwDH0/p2817333930.webp |
| 15   | 狙击手                                                       | https://i.postimg.cc/JnzCnmGT/p2738601191.webp |
| 16   | 奇迹·笨小孩                                                  | https://i.postimg.cc/3Jm8z8pp/p2842327103.webp |
| 17   | 长津湖之水门桥                                               | https://i.postimg.cc/8PQDR23C/p2846021991.webp |
| 18   | 这个杀手不太冷静                                             | https://i.postimg.cc/g0TbnQ8m/p2831482222.webp |
| 19   | 熊出没·重返地球                                              | https://i.postimg.cc/BQBWxXHh/p2856825681.webp |
| 20   | 喜羊羊与灰太狼之筐出未来                                     | https://i.postimg.cc/4yv0bS7K/p2797468943.webp |
| 21   | 小虎墩大英雄                                                 | https://i.postimg.cc/vHzj4Cwk/p2792787666.webp |
| 22   | 十年一品温如言                                               | https://i.postimg.cc/8CWm6rrc/p2734149986.webp |
| 23   | 好想去你的世界爱你                                           | https://i.postimg.cc/2jJyP9tp/p2686054168.webp |
| 24   | 京北的我们                                                   | https://i.postimg.cc/fLyWPYjT/p2857189540.webp |
| 25   | 不要忘记我爱你                                               | https://i.postimg.cc/wxSJhBgs/p2842142874.webp |
| 26   | 尼罗河上的惨案                                               | https://i.postimg.cc/tRVnCHWR/p2854856216.webp |
| 27   | 我们的冬奥                                                   | https://i.postimg.cc/bYbdf9fs/p2828682728.webp |
| 28   | 纽约的一个雨天                                               | https://i.postimg.cc/9QxQYxZN/p2856549164.webp |
| 29   | 隐入尘烟                                                     | https://i.postimg.cc/k4SM9qZk/p2855320664.webp |
| 30   | 记忆                                                         | https://i.postimg.cc/zvSJbpq7/p2666266209.webp |
| 31   | 人间世                                                       | https://i.postimg.cc/MpGWHyN0/p2764924182.webp |
| 32   | 卧鼠藏虫                                                     | https://i.postimg.cc/NMRtSb2t/p2782052522.webp |
| 33   | 向着明亮那方                                                 | https://i.postimg.cc/K8VyYCnF/p2831434654.webp |
| 34   | 跨过鸭绿江                                                   | https://i.postimg.cc/63xJrRjm/p2772730001.webp |
| 35   | 萌鸡小队：萌闯新世界                                         | https://i.postimg.cc/GtzWnfcW/p2781907425.webp |
| 36   | 永恒族                                                       | https://i.postimg.cc/Fzcxjd6D/p2677303737.jpg  |
| 37   | 黑客帝国：矩阵重启 The Matrix Resurrections                  | https://i.postimg.cc/rwMvT2Vt/1.webp           |
| 38   | 法比安                                                       | https://i.postimg.cc/hPcmC5QS/p2686474863.jpg  |
| 39   | 命运/冠位指定 终局特异点 冠位时间神殿所罗门 Fate/Grand Order | https://i.postimg.cc/VsRvNL58/p2648204312.jpg  |
| 40   | 精灵旅社4：变身大冒险                                        | https://i.postimg.cc/y655c0DX/p2659301260.jpg  |
| 41   | 杰伊·比姆                                                    | https://i.postimg.cc/2yD1K2TR/p2734251152.jpg  |
| 42   | 法兰西特派                                                   | https://i.postimg.cc/44679RNJ/p2634539726.jpg  |
| 43   | 世界上最糟糕的人                                             | https://i.postimg.cc/gjm0DrkG/p2669034145.jpg  |
| 44   | 驾驶我的车                                                   | https://i.postimg.cc/DfP0zQB0/p2639821491.jpg  |
| 45   | 天鹅挽歌                                                     | https://i.postimg.cc/dVptf5tH/p2717809625.jpg  |
| 46   | 健听女孩                                                     | https://i.postimg.cc/BZpZ3Sn1/p2665001891.jpg  |
| 47   | 超能敢死队                                                   | https://i.postimg.cc/9FfmF6Dg/p2685536675.jpg  |
| 48   | 麦克白的悲剧                                                 | https://i.postimg.cc/YC5MRjbT/p2722980045.jpg  |
| 49   | 圣母                                                         | https://i.postimg.cc/3w27c7qN/p2713726815.jpg  |
| 50   | 无辜者                                                       | https://i.postimg.cc/c1FWjFFR/p2669238453.jpg  |
| 51   | 不要抬头                                                     | https://i.postimg.cc/fWGsZGfN/p2730833093.jpg  |
| 52   | 古董局中局                                                   | https://i.postimg.cc/tg4Q2pph/p2734316987.jpg  |
| 53   | 欢乐好声音2                                                  | https://i.postimg.cc/vTRKFsF2/p2732782860.jpg  |
| 54   | 魔法满屋                                                     | https://i.postimg.cc/SsqG4N1P/p2807936075.webp |
| 55   | 沙丘                                                         | https://i.postimg.cc/T2r7LBXG/p2687443734.jpg  |
| 56   | 犬之力                                                       | https://i.postimg.cc/hv65nQzB/p2678298618.jpg  |
| 57   | 倒数时刻                                                     | https://i.postimg.cc/QdLyn5Zg/p2690712224.jpg  |
| 58   | 铁道英雄                                                     | https://i.postimg.cc/XJvHyk9W/p2684720964.jpg  |
| 59   | 花束般的恋爱                                                 | https://i.postimg.cc/XqGssN2D/p2623936924.jpg  |
| 60   | 芬奇                                                         | https://i.postimg.cc/jqDTpzmB/p2721066869.jpg  |
| 61   | 偶然与想象                                                   | https://i.postimg.cc/9X66ssGq/p2687984714.jpg  |
| 62   | 长津湖                                                       | https://i.postimg.cc/T39BLwBr/p2681329386.jpg  |
| 63   | 最后的决斗                                                   | https://i.postimg.cc/R0p8dYSh/p2672789902.jpg  |
| 64   | 上帝之手                                                     | https://i.postimg.cc/vTfNQHWx/p2688141684.jpg  |
| 65   | 兹山鱼谱                                                     | https://i.postimg.cc/mDFdw2pp/p2634952893.jpg  |
| 66   | 六号车厢                                                     | https://i.postimg.cc/kDr4pCyP/p2702495844.jpg  |
| 67   | 毒液2                                                        | https://i.postimg.cc/kGS84f19/p2675102928.jpg  |
| 68   | 007:无暇赴死                                                 | https://i.postimg.cc/g0B6Hf1m/p2707553644.jpg  |
| 69   | 浅草小子                                                     | https://i.postimg.cc/VLc0g9Zr/p2675919720.jpg  |
| 70   | 暗处的女儿                                                   | https://i.postimg.cc/bwmGq7yk/p2807771124.jpg  |
| 71   | 希尔达与山丘之王                                             | https://i.postimg.cc/Bvtjbmby/p2799276725.jpg  |
| 72   | 黑寡妇                                                       | https://i.postimg.cc/ZRd0F0rx/p2665872718.jpg  |
| 73   | 记忆                                                         | https://i.postimg.cc/2yd6MJ9q/p2666266209.jpg  |
| 74   | 慕尼黑:战争边缘                                              | https://i.postimg.cc/Wbnz7BJC/p2696233262.jpg  |
| 75   | 小人物                                                       | https://i.postimg.cc/t4vgKFkz/p2640615589.jpg  |
| 76   | 斯宾塞                                                       | https://i.postimg.cc/HLJLdjP1/p2678164205.jpg  |
| 77   | X特遣队:全员集结                                             | https://i.postimg.cc/d3FQ26QJ/p2637084112.jpg  |
| 78   | 新生化危机                                                   | https://i.postimg.cc/XqgjQ5Xd/p2692391480.jpg  |
| 79   | Soho区惊魂夜                                                 | https://i.postimg.cc/kMNgn0wG/p2704450439.jpg  |
| 80   | 天赐灵机                                                     | https://i.postimg.cc/VLDstXxc/p2681332523.jpg  |
| 81   | 尚气与十环传奇                                               | https://i.postimg.cc/Dz1trTSb/p2674321872.jpg  |
| 82   | 呼朋引伴                                                     | https://i.postimg.cc/DfxRd5w5/p2828740205.jpg  |
| 83   | 私人荒漠                                                     | https://i.postimg.cc/hvvy5Wts/p2699152497.jpg  |
| 84   | 新网球王子 冰帝 vs 立海 Game of Future                       | https://i.postimg.cc/wx4bsvwG/p2619484634.jpg  |
| 85   | 正发生                                                       | https://i.postimg.cc/fT3q7NFT/p2681278523.jpg  |
| 86   | 开端                     | https://i.postimg.cc/8C7wpzGk/p2817285601.jpg  |
| 87   | 今生有你                 | https://i.postimg.cc/htj1stf8/p2853101640.jpg  |
| 88   | 青春训练班               | https://i.postimg.cc/XJMKnKxx/p2853484859.jpg  |
| 89   | 废柴的一日三餐           | https://i.postimg.cc/wvhXMtGz/p2745531024.jpg  |
| 90   | 辛普森一家 第三十三季    | https://i.postimg.cc/fTg9wjRX/p2687674570.jpg  |
| 91   | 和平使者                 | https://i.postimg.cc/4xDHcftx/p2759888858.jpg  |
| 92   | 勿言推理                 | https://i.postimg.cc/j5gszc94/p2841026061.jpg  |
| 93   | 超越                     | https://i.postimg.cc/XNKMSQxH/p2829656990.jpg  |
| 94   | 流光之城                 | https://i.postimg.cc/rmSHTwG0/p2837881351.jpg  |
| 95   | 对手                     | https://i.postimg.cc/Y0MsCB8X/p2782204672.jpg  |
| 96   | 守护解放西3              | https://i.postimg.cc/rw0QcpMV/p2808177016.jpg  |
| 97   | 鬼灭之刃 游郭篇          | https://i.postimg.cc/5y2Sr4sM/p2851359433.jpg  |
| 98   | 画江湖之不良人5          | https://i.postimg.cc/BQ8xPwyN/p2797706656.jpg  |
| 99   | 一年一度喜剧大赛         | https://i.postimg.cc/CxGnjy4C/p2701105839.jpg  |
| 100  | 双峰 第三季              | https://i.postimg.cc/13hJmQtD/p2434434204.jpg  |
| 101  | 亢奋 第二季              | https://i.postimg.cc/TYCgSvbR/p2806778258.jpg  |
| 102  | 淘金                     | https://i.postimg.cc/G2xHyvJc/p2855160940.jpg  |
| 103  | 那年,我们的夏天          | https://i.postimg.cc/zvdvd5c6/p2741745122.jpg  |
| 104  | 镜·双城                  | https://i.postimg.cc/QthMhHLv/p2838047633.jpg  |
| 105  | 但是还有书籍 第二季      | https://i.postimg.cc/0jH5Yn6q/p2629357934.jpg  |
| 106  | 雪中悍刀行               | https://i.postimg.cc/R0zSLN0q/p2714598385.jpg  |
| 107  | 大妈的世界               | https://i.postimg.cc/LsBHC2qD/p2812282180.jpg  |
| 108  | 爱很美味                 | https://i.postimg.cc/Vv5mktKZ/p2747567321.jpg  |
| 109  | 进击的巨人 最终季 Part.2 | https://i.postimg.cc/bN7pqskL/p2728015364.jpg  |
| 110  | 国王排名                 | https://i.postimg.cc/W4nvm8Kf/p2681362557.jpg  |
| 111  | 王牌部队                 | https://i.postimg.cc/bYMjRsC8/p2623510175.jpg  |
| 112  | 契×约—危险的搭档         | https://i.postimg.cc/gJT14TK4/p2734589915.jpg  |
| 113  | 家族荣誉                 | https://i.postimg.cc/6phWzB40/p2826837678.jpg  |
| 114  | 雪滴花                   | https://i.postimg.cc/857T1WKK/p2722034745.jpg  |
| 115  | 成瘾剂量                 | https://i.postimg.cc/2yvmWYfn/p2682534549.jpg  |
| 116  | 瓦尼塔斯的笔记 第二季    | https://i.postimg.cc/vHdXnVWX/p2773991350.jpg  |
| 117  | #居酒屋新干线            | https://i.postimg.cc/kXcNthMX/p2765302469.jpg  |
| 118  | 安检 第二季              | https://i.postimg.cc/N0WRw3Db/p2836015917.jpg  |
| 119  | 沉默的真相               | https://i.postimg.cc/q7CCQmrK/p2620780603.jpg  |
| 120  | 妻子变成小学生           | https://i.postimg.cc/x89JfnzG/p2790697478.jpg  |
| 121  | 解读恶之心的人们         | https://i.postimg.cc/1zp3fbPv/p2828771329.jpg  |
| 122  | 半熟恋人                 | https://i.postimg.cc/W1Z2HxDs/p2808055043.jpg  |
| 123  | 觉醒年代                 | https://i.postimg.cc/504bj0st/p2631873666.jpg  |
| 124  | 完美伴侣                 | https://i.postimg.cc/7YGD5jnS/p2647059010.jpg  |
| 125  | 英雄联盟:双城之战 第一季 | https://i.postimg.cc/jd70xPRX/p2714077426.jpg  |
| 126  | 一闪一闪亮晶晶           | https://i.postimg.cc/DZZHtbC1/p2842225770.jpg  |
| 127  | 酒鬼都市女人们           | https://i.postimg.cc/s1FLQH3J/p2681862170.jpg  |
| 128  | 请回答1998               | https://i.postimg.cc/DykFbBjD/p2272563445.jpg  |
| 129  | 风味人间 第三季          | https://i.postimg.cc/DzWY6Kmc/p2765048624.jpg  |
| 130  | 81号档案                 | https://i.postimg.cc/Lsf3h4nG/p2791066497.jpg  |
| 131  | 沉睡花园                 | https://i.postimg.cc/nzGKYdKk/p2768730564.jpg  |
| 132  | 你是我的城池营垒         | https://i.postimg.cc/rsy1hSQL/p2635182901.jpg  |
| 133  | 他不是我                 | https://i.postimg.cc/nrLmKnhF/p2752910229.jpg  |
| 134  | 琅琊榜                   | https://i.postimg.cc/WtL2N7KB/p2271982968.jpg  |
| 135  | 御赐小仵作               | https://i.postimg.cc/Bb1Sfzh5/p2652208493.jpg  |
| 136  | 飞屋环游记               | https://i.postimg.cc/XYLqVVG7/p2363116942.webp |
| 137  | 加勒比海盗               | https://i.postimg.cc/W4pBvVgV/p1596085504.webp |
| 138  | 贫民窟的百万富翁         | https://i.postimg.cc/fThgsZQP/p2434249040.webp |
| 139  | 荒蛮故事                 | https://i.postimg.cc/c4mTNnmK/p2584519452.webp |
| 140  | 恋恋笔记本               | https://i.postimg.cc/RVX1YzSK/p483604864.webp  |
| 141  | 少年派的奇幻漂流         | https://i.postimg.cc/fLN0LHWM/p1784592701.webp |
| 142  | 辩护人                   | https://i.postimg.cc/6p47hsy9/p2158166535.webp |
| 143  | 疯狂的麦克斯4：狂暴之路  | https://i.postimg.cc/YCF9W1s0/p2236181653.webp |
| 144  | 终结者2：审判日          | https://i.postimg.cc/wxt7z8QD/p1910909085.webp |
| 145  | 小森林 夏秋篇            | https://i.postimg.cc/QdqMmXjZ/p2564498893.webp |
| 146  | 神偷奶爸                 | https://i.postimg.cc/JnS7005Z/p792776858.webp  |
| 147  | 聚焦                     | https://i.postimg.cc/mZF2q3Tw/p2263822658.webp |