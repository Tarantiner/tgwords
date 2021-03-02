package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

var wordLis []string

var err error

func init() {
	// 加载词库
	var f *os.File
	f, err = os.Open("../sense.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dd := json.NewDecoder(f)
	err = dd.Decode(&wordLis)
	if err != nil {
		panic(err)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindWords()
	}
}

func FindWords() {
	s := `
你公司还要人吗现在没路子这图带字的谁白熊嗯我最近没联系他哥们你不看嘴型‘？下午好不好吃算了吧这个群还要吵？大半夜的收波人特区接3万内赔付认准王祖蓝另收个英文翻译大佬�有意向来私聊呀本人照片差距大不1一起回家好好说这个狗币乱打的，，，而且我俩之前在一起我睡着的时候还拿我手机用我的支付宝充值去把自己身份改成本地人
就可以自由出入了@alang112233狗犊子你为啥不理我什么事儿画画的北北画画的北北王者荣耀吗秒瓦滴出四件套，工商。农业今晚又要去？泳池十点红莞会所A牌原价998M牌原价1498Q牌原价1698T牌原价1898S牌原价2198进门说找（红牛）就算预约买单A牌立减100其他牌立减200全套莞式水墨90分钟不限次（主打中国妹子）不报（红牛）买原单本店提供
免费饮品，车接车送也可以送上门营业时间：老挝时间19点至凌晨06点白天值班（红牛专属）老挝时间11点至下午3点VX：a2772800438app：2095873064飞机✈️群：@jinsanjiao888888你是方便微信还是支付宝谁说没隐卖奶茶那家的遇见各位大神好�老朴在哪里？���金三角没有爱情再也不敢在这飞机上找商家的。太不靠谱了像我老家有没有跑腿的
卧槽群主挂逼了我的宠物好感动哦，人家都快哭了这个不叫人妖，叫变性人网恋选我我超甜又骗感情又骗钱办公室坐着了养老的四件套现货，现在喝芋泥啵啵奶茶的都只要波波了1没注意看这里也有你特区没女的了吗男的你们都这么感兴趣我也有这娘们的微信急！特区出租车！刚才一名白色t女性从蓝盾侧门到华源酒店黑色手机忘后排了你把他们全拉
来做广告赞助看没龙湾里有新的少哈哈哈你好骚啊�骗子你叫那个华仔特区哪里有修屏幕的手上有现金私聊我我俩不是一个会所得我还能说瞎话？早点没关门呢吧，超市那种也不是药用的缅甸都要差不多一晚啦知道吗好兄弟这有啥卧槽不是我卧槽一颗牙齿两千？？？？？？？？打架吗谁扛得住偷渡过来的恰了3岁？good我都要我微信有个广西妹子我要
上号去了1钱记得给仔细一下，没胸@fengzi9989封城都是小事3万的赔付还说最多接多少赔付正规的我在啊微信推一下你这工资太低了他胖我都排人不上号黑不溜秋的怎么下口�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主你不讲武德出特区把特区的四千给了然后到国门给国
门的一万是吗到时候带我回家特码哦看中什么了那把妞这个字给我去了吧不符合我的气质六合就靠自己没期买一点头一次发现啊到地方之后报备又不花钱会所来一盒今天是临时起意的不交钱还这么硬气打广告好奇怎么样找我我也赚不到钱刚刚删了我刚要按。我才来啊有没有人晚上打篮球的好的我不敢商机签到����？？？可以啊都怪你每天私聊啊
。操怎么去麻痹！好眼熟在床上葛优躺没护照少打点炮，每700人就有一人有艾滋特区也贵啊你又不照顾我你们难道不发短信推广码？找我发国际短信，来的都是首充百万的��我不要求多，一个月10万让我赚就好回家？不存在的偷渡并不是大罪明年年中回去怎么样。来吧退群要不要我送你不睡觉嫖男仔啊天天开房她只会送你回去回去肯定都要去报道
下吧输死了打了牌搓女人不贵聊到现在为啥要回国麻烦推一个，谢谢法克？桃总有没有牛肉汤的微信吗？谁有小李批发部微信？缅币吗？其實雞巴大不大無所謂最重要是你喜歡的樣子高了好的哦男的女的发。我看广告广告你长的好像狗蛋清什么卖代步车自动挡的号商看到私我号商看到私我号商看到私我晚上带来给你看1让跑腿送来我给跑腿费你也可以
做我出2000可乐有便宜的被可乐气瞎了大曾找不到一个微信吗开了那段时间很慢包夜滴滴安排甜蜜!!�有人去吗好吧可乐不在了那不行还没有妹子看上我不能退群狠到了艾特她缺个女朋友这个暴躁的樱桃我有空哪位大佬可以办出入特区通行证你还插在手机上这说的是真话签到你错了，是niuroubing8888只能摸没电没网能发信息吗？算啦不问了800！喜
欢滴滴！他们会爱上我的？怕女管理看了受不了喝个屁保护费交了吗就打广告跑腿的微信有吗拉你进群了不多比比罚款了可以直接走吗养肥了在杀你爽了白干就白干啊考拉是美女没有就睡了嘛弟娃这里是不是波贝？这抖得好厉害对有想吃烧烤的老板吗有没有美女✨�华侨国际会所�✨�A牌预约价1298¥90分钟服务�T牌预约价1598¥90分钟服务�M牌预
约价1998¥90分钟服务�W牌预约价2398¥90分钟服务�S牌预约价2698¥90分钟服务�（豹子号）预约价309890分钟服务SS日韩牌3198¥70分钟服务��️买两个钟包夜6小时?  ����进进店报@明明‼️预约@明明‼️优惠多多‼️�做爱（最少八招声情并茂）出水不限次数✅✅✅�服务内容:本会所所有妹子都是全套水床服务，均不限次数，制服，艳舞，
袜，胸推，毛扫，臀推，深喉，口爆，毒龙，會員充值1.0.0.0.0送1798积分充值2.0.0.0.0送两个1798积分充值3.0.0.0.0送三个1798积分（日韩除外）可以选任何技师，积分可以累积叠加使用讓您高興而來，滿意而歸，�24小时营业���每天20位以上美女金三角秒安排欢迎您的光临��������本人诚信接单专车包接送‼️纸飞机85620521171
78微信号18509265951WhatsAPP2052117178啥药店的飞机来个我是在菲公司老挝正规入境办理一个车上就坐我一个男的其他全是女的应该总会有一个女的愿意的就是这个人摸奶我怕她掏出来比我的都大可以吗我有个卵啊如果没中我跟渣男买了老板娘买了嗯哈哈哈之前9码想错了，再想给我地址我也去没有，你退群吧人呢来的时候120，现在168去去去，
我可不喜欢你分妹子你要不要租个房子@hzz9999早安打工人。你才发现给5块我刚到特区落地负债50w，起步怎么能放弃呢？皇上驾到你问下这个VPN大把的，就看好不好用了有没有要渣男陪睡的300包夜！！！明晚绝对我说的这十码出你又不是政府高官18cm不含头叫你去？禽兽我来咯是啊宝这样的行不问了好多遍了滴滴我就在他旁边坐着你又知道了你
吃别人的东西说不好谁有鸡儿图来几张…有菠萝蜜吗到底多少钱有别人吗马上要换网了吧你是小鸡饽饽？我们中午休息两个小时晚上休息两个小时。你们13小时呀宝宝心里苦谁跟你说我没对象的大曾破费了我好歹也是副业拉皮条啊没事的我不兼职没得事我知道美食街我找人立正产业园吗正规足浴那里有带着护照忘记签了我大概亏，几百块财神酒楼的
微信谁有呀啊哎，又一个小伙子被骗了二百块钱他不配我踢我打字吵架真没输过来现实吧我答应了一下没有乐趣了我都来好几天了坚持不懈哈哈哈不要钱这还不便宜日本妹子特点虎牙和兔牙在木牌很好找女朋友哈哈哈哈听说产业园si人了下班回去鞋子穿起来在哪里哎哟干嘛这么低调老板不能换来帮我吵群吧，我也是吵群的优秀淦到底是个什么梗算了
我不说他了优秀去吧来一首我的唇吻不到我爱的人天天洗脑约吗？全是卖的是车牌号不知道有没有看错小姐姐你在哪家酒店热哈哈哈哈这图怎么看都在像是说我草泥马�刚还说吃腻了我不染发不烫头纹身在里面也没见小姐姐看上我啊哈哈哈不需要啊切肥狗群主这样子买了金三角哦哦，或者晚上去我宿舍聊天借我聊一个你们发工资了吗的确不像没有不
上头阿哈哈哈looklook逼？多简单吗唐人街咯记录一下猪腰子鞋谁有中国电话卡联系方式你在哪家喝香蕉牛奶果汁字母容易打偏她发过伊利没有大屌萌妹上个星期我三个朋友回去了然后又说回去提前告诉他健康码可买梅龙汪，有护照的会不会有人身安全问题我随时随地都可以出去少妇喜欢小鲜肉还是告辞我来给你温暖只是见多了哈哈三十岁的女人如
狼时候广告不要一直发问我啥招男模吗体会过棒子与棒子之间的摩擦吗我也是喜子尼玛赚的比杀猪的都多去迪拜秒回的我你喝了一个小时了怕这怕那在家待着好了哎这是在老挝人家棚子里吃的鸡汤红馆会所群算给个面了爬山来的不是免赔付的吗感兴趣不有赔付吗你比较需要回收二手老娘们，不锈钢盆换我公司上个杀了一千都是小姐姐越南女的梅陇湾
铁锅炖你荣耀王者15⭐的怕啥LOL吃鸡有组队的吗这个是在哪里参加的我觉得还是富婆比较好你们吹慢慢扯狗儿子确定没有想起来吗是人才留着自己用不香嘛什么条件怎么说可以阿报啥特区第一人，嫖娼吃伟哥~他们不是尿关注这个机器人了吗舔就完事儿了没有，退群吧�嗨皮去了这边有健身房嘛？我天天在这唠来告诉我她的联系方式，你这个不是拉
皮条，你这个是为拯救失足少女行动做出伟大贡献我开的一千多木棉酒店有个单间大床房要的联系我不带有护照750到磨丁没护照1W左右看你自己找的渠道了有没有奶茶店我上次在我老乡店里给妹子按摩还收了五块钱小费呢有药店的微信吗？t11什么时候上班呢？我都是借的不甜有没有牛皮的拉手，联系我。特区外的外面来一个准备好去你家找妹妹了
你帮我接，到付首先先搞定他小弟没呢除了日本好像都有只是多和少的区别女人还可以不过现在好了嘎哈知道他在产业园但是就是不露面我工作机都下载了群里有没做美发的�什么狗我在问大老婆双胞胎能一起吗？一起啥价格差那么点，我总不可能自己掏少喝酒，文明开车因为没有工作，找不到公司报备呢这少妇我喜欢请就推丢来个qp大佬好渠道助
你业绩超级加倍我管你什么没有吧，我记得楼下就是再说下去就剩四个半小时贵圈这么乱八戒操计划泡汤了没阿滚过一个关口就要被查一次有人在找台子吗。太少电动车�有人卖吗卖不锈钢脸盆哪家有卖过桥米线的吗星耀4西港会所还便宜发照片啊掉4颗星@toudanquan在不缴费，小心给你踢了主要是紧20有快送前所未有的精神别说兄弟不是人有没有卖
港卡的卡商�我tm我还以为你是我姐我给陪酒妹整上头了别把自己当个藕认真的吗？我是下面厉害不认识有啊和白熊回不起啊有不想回家，找公司的小伙伴吗���看看美国大片都知道口活的好像不是会所的私人接单的拉皮条去了我把我上个月赚的全给蓝盾报位置我怎么教你？一个粥48呢禁足了？mmp人家不给不要扯蛋牛逼的中老两文的接了我要求吃
喝拉撒都在园区，不能出去这就是你的不对了我在家每天最少10小时睡觉怕啥，人美啥都不用摸无护照还还办理什么什么系找台子谁拉我进两个片群甩个女孩子现在没有小姐姐医院联系方式这俩者完全不一样的性质你懂吗？几个点正常个屁啊，在国内就不正常，在这边就正常？他打鱼艹话粗理不粗你发的这是太子？我不是不回你你这么可爱感觉回什
么都配不上你把他埋了速度呢你说个几把必须便宜喇叭声音开到最大群主她妈红塔山了只有小黄楼有妈的美？？带上你的小音响私聊谈谈她卖自己都买不到我女朋友我有空啊黄色的。5字头你安全回国就好啥的关心了我一波告诉他在哪里在哪个公司，我去恳求他，下次不要这么做了为什么哈哈哈哈哈来多长时间@wb123789没有好的身材很好那个的头像
啊我说你都不会看我打的字？群也封了又不能玩还是有实力的妈的线下会员三百中了个码群里有需要引，流的吗���找工作私聊我~辩证的角度看事务吧，宣传是那样，所以你就只会那样说你说了什么话！被拘留了感动万象隔离过来在特区外面还要隔离?认不到/addnewchat@sauweenbot我喜欢中了就没有买就今天啊你是天天在会所的路上赶紧发啊录
个视频呗就怕烂写盘，各种台子�️�️就很难受我想多赚点钱养他然后他不要回国报备没有什么影响吧。可以出不可以进吧如果工作不开心，告诉我，我上门接你一码中特你想看什么照片纸飞机在国内是不是就用不了了一看就知道身材也很瘦，摸着不舒服。。都看不见你这么激动现在还好不收男的我又不是拉皮条的刺激我干嘛那么眼熟我要半个月
后再去可乐我们和小樱桃也要吐我长得丑，不敢上去搭讪出一手QQ靓号，Q解冻有没有美食城成都串串的微信为什么哎，我特么的好累给2500，11座的就接你一个。我退了猪肉铺咋了做菠菜不安全你明天几点换班啊来嘛，红磨坊会所这里永远一直在等你们。特区外面东盟是不宵了我是个帅哥你喝多了没上人家？真的假的????????那么我走我也要小尾巴
我有点小意见了我怕被强奸去不去？晚上给你安排几把2万一个晚上出去你好他天天口嗨说唐人街结果让他出去了他说怕我们绑架他？？？我朋友有鸡店让他安排妥当即可啥意思啊不准我们去？老板醒了@yfsg888这么多这个群你不要发水果嘛不可以哦他不是女的害有没有好吃的炒饭外卖推荐不去要饿死也不是不可以我们一起去在加饭鸭脖怎么卖的要不
是我忙三个大长腿妹纸啊群里找老能遇见你送我出去被一般我只想好好吃饭渣男在我这他们不配有名字。大曾说的我不知道给大曾甩一个吧省得大曾天天晚上夜班打飞机也没叫我你昨天晚上发我们照片了吗但是你送不进来渣男呢？昨天晚上你们比帅她妈的鸡吧渣男别拉我牛逼啦特区不是不让电嘛蓝牙耳机吗私信我一个300买一送一我们有缘吗拿过来我
会你这是玩小妹妹玩累了才来这里吹牛逼的吖群主在哪裡是啊吨吨吨我在9栋好的谢谢了那你怕什么狗曾是老年人没有睡眠的我才起来你就要干饭了他妈的吃了悍马糖睡不着了哈哈���为啥我前天去没什么上个锤子锤子不去不可以卖肉的呢虽然说在特区很多出来玩都是做鸡的水果妹300斤放个屁都能震是你们不说踢了那些越南妹纸在木棉酒店干嘛不
会的，它不会对蚂蚁出手的我的肚子发炎了，妈的玩过了哈哈不然？不知道啊还好不多，只骗了我一千多块钱。令狐小生15859667621云南省西双版纳傣族自治州勐腊县勐腊镇勐腊新城中通快递吴小兵转金木棉中国快递刚刚不知道哪个�的乱发语音属于性骚扰没有要听妹纸讲话吗代理了解下流水抽成还有z工资长夜漫漫那个二维码不全看吧这不就好了
可以的好无聊@baofu8888残疾人？买什么人有没有卖中国手机卡的特么的都是小白好吗你说到了，菲到了。��伊利今晚要开房了认真的叫你给我上个管理，你很啰嗦哦还有去蓝盾的去哪里干嘛是啊小贱有没人谈恋爱的那种谢谢小哥哥们今天会有六合吧没把她当女人上班这么早对啊起床穿衣服了不想回去哈哈不知道咋了哈哈哈快发你的图片�欢迎加
入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主报备过是去会晒吗狗曾呢四件套是啥找死？特区内有没有黑卡卖？大量出优质gm米分老挝啤�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主贼香对不起……我看
成了花圈是啊走了会怎么样@wb123789让人认识你好家伙也人称老挝通还贼安全都没有保障没钱啊待会下班给你看个好东西，私发姐妹们你们争继续你们的表演丰满一点的才有感觉啪啪啪才有声音可以尽情的去打炮�️�银行卡4件套确定，因为我才是一手你又不是我的宝贝客户签名这就是为什么我皮肤好的原因垃圾这都快过年了抚州离得近一个人吃
吗是不是经常约你出去吃饭喝酒啥的我还没洗…继续我被欺骗了好晚上就去�你有多少卖给我250011有mmp？颜值不够画画来凑。你背叛主治难道分手了就不能做朋友吗不是这家握草干哈你还热乎的等下再来问我要是回去，要找广东商会给我冠名一个合法身份我才回去上管理了吗你唱不信博翔刘德华没见过吧也许吧就是下不来啊，他妈的这边要收费的
做人不能那么贪心签到我在博翔咋可能是你想去就能去的。洗脸修眉你在美个甲不可能我今天截图了去呀肯定啦不能发大量大额云闪付码懂的来收张手机卡有的联系我刚刚的视频太过刺激你他妈的逼脸呢祝你落地三件套4000.8000�骗你妹换了钓凯子行吧我看看直接打死金三角没有女人了吗？打几把下午好小推推欢颠没见过是的，不过有的身材好又好
看也就500泰铢缅甸的太黑了分分钟搞定金三角第一贱接人就出来了。。。昨晚那个美女呢可以两个人一间现在辉过报备得几天有人知道不你吃饭的时候注意看哈哈哈有点273747绝壁7尾上期开的25收个二手苹果x手机，有要卖的私比酒店爽傻逼，你要男的？签到那我可以白天你你以为你忙完了应该人就来了嘛我在上班看手你要吃嘛继续说一下别太想哥
了？怎么回事的老铁夏天阿姨你收几天了还没收到没有东西送昨晚快乐吗是走国门了以后还能出国吗？2100签到都叫宝贝快点直接去呀你咋样尼玛的老挝泰国的我不知道啊昨晚浅深不错哪里的管用吗他妈的都不用上班嘛你是怕我榨干你吗嗯等你我先包礼物我不配他妈的发了半天他都没出来你是中国的吗。我也要吃你要吗不来？两万能接么等下被投诉
哦还打雷不正常马上睡午觉了二十号解除带不带家属？先回国我才15岁你可以云储存啊我们办公室就老大有马子一样二十四小时有人上班黄牌，红牌我都有天天你是不是不睡觉舔狼是单钟就你能群主你可以退了抽抽抽我这样去嫖娼小姐会不会不给嫖啊下次去韩国花点钱给她留下我的基因泡面都没得吃的怎么说野战不安全我才发现我有这个微信，我记
得我删了刚看还有红色的要好点有爱国汤吗你喜欢我？所以查手机的别说了兄弟们你没有管理啊正规的塞林木哪里收赌场筹码我这边需要一个养微信号的人一个人的好天气舔你一嘴的尖锐湿疣我怀疑你是70后拉皮条的这里酒吧也不算贵吧。。。千年老树开花了飞机还是微信我现在还不知道怎么过去感谢一句话都没有是的你喜欢闻狐臭？这癖好第一次
见想玩用的电路应该不一样好的过来拿我穿着旺仔牛奶套装我不私聊再见早安大家跟他开玩笑�我没有哦可能是我喝了酒的原因没那个体力干了一小时三次特区花硬币咋了？很多？你点十个我都给你点卖男装的有没有私一下真的假的收费太贵了�签名有没有中国卡现货的妹子都去做名媛拼丝袜了���你喜欢有没有女朋友找一个冬天来了了谁有汇丰
酒店微信啊乖孙我不知道啊你去群里面看看炒群啊哈哈.杀猪区块链国内外都有赔付是不好听的，转让费不好吗吃了补肾�来开卡丁车走请你吃麻辣烫多少钱我有疯子哈哈哈说什么你们不会乱发吧，名字都没写黑心号商那天状态不好然后全程手机响不停唉帮我置顶一下属实没点毛病我不玩游戏的大变态手机如何翻墙回国啊两个加起来我做饭好吃私聊特
么的勾引客户碳元素好吗群主把你们公司的妹子拉进来好吗白天有没有小姐玩到群里发情来这里有多少给我来多少个另外一个号啊你猜为什么不在国内寄呢我上班12小时纸飞机在线10小时不逼逼了不是还有好多国内警察在红牛有掉用，给我悍马糖刚上班在这准备炸群呢业绩就是尊严嗯那可能是我记错了协议拉40人群接一手二手数据导粉精准定位数据
可洗（棋牌股民彩票金融）需要老板私聊我资源很多罚这么多打个小卡吃个鸡儿开车啦还是免费的你我弟兄，辈子弟兄1狗推我们一起去淋雨牛逼搞笑的没事吧电影名称，丧失夺命大奶子这次不一样金三角欢迎你老板都好骗，你说小妹能不好,,Ծ^Ծ骗吗。你给我介绍一个客户我不就开单了你馋那就吃啊哪个傻逼拉的我云贵口味，做工作餐的联系下椒房
楼上是什么样的啊大佬卧槽有个傻逼跟我聊骚签到佛系接单，拒绝一切理由白嫖艾特我是想我了嘛有隔离14天，然后交点罚款就完事了这么嚣张哈哈哈我有8盒哈哈哈哈你想怎么谈宝贝那私发给我十点半差不多漂亮疯了啊你又不回去有视频吗jb打飞机都打肿了Papapa♥️♥️♥️♥️♥️什么叫广告我直接一次满足你我草
你哪里是什么地方群主内裤要不要？管的太严了。这里卖鞋的，质量好差�中文频道去看吧你们这些人心思坏的很给你了权限这个马甲不就是别人的了吗私聊听说黑的一比睡觉人美，逼美我想跟你在一个房间隔离十五天啧啧3400不能买东西？好用的电脑多少一台你这个360度大转弯啊晚上都来报名吧充的有点多不要用啊干不了1小时，我倒找你200我没
有玩不要你是男生你不会主动点吗砍妖秀你男的？我见过啊老乡啊17年来的我们难道就没有风险吗人家还唱歌呢没有啊你的什么药？就华侨贵点只支持国产嘎嘎香的那种这小尾巴可以哪里有卖洋酒的？包夜500的那种你下去鸡巴不群主叫了我现在就过去找你啊哈哈哈哈对啊���多少钱？中国鱼我到了，下来拿吧省钱给老婆话完事了，拉走了早应该有
我想知道你们养了不杀嘛？那你们吃土啊你是这样杀猪的？.......都是骗子妹子呢？不是兄弟往上翻，呻吟超好听还是你的原味丝袜©47啊高请使用文明用语，你已被禁言1分钟还有10分钟下班了md给力害那你还有机会鸡巴，要就来拿，叽叽歪歪一堆有屁用哦那我呢对吖隐私了好刺激哈哈可以！晚上約一下！一起打遊戲不能用碎片换吗？一月保底让我
赚多少安溪的吗？早1有没有打炮的能够知道你是整过公司上班的就挺好的啊噗....女孩子啊哈哈哈哈签到没有收到我们真刀真枪的干，你就知道我的厉害宝贝消费这么样谁有汽修厂的联系方式，麻烦给一下！谢谢你吃饭，买东西，开酒店需要回国网络的联系我别走，有钱有钱哪里理发便宜啥情况？大佬们好这个看了心疼有多高这么大个群，你们打篮
球在哪里打牌好烦啊，无聊的很。有没有妹子，来一场强奸大戏不，我只是我家月月的宝贝你来了没有嘛我是给你说方向十分感谢群主怎么当的来个玩彩的大佬不限IP哦谈恋爱吗？哪位大哥知道二字头的卡刚买来怎么使用有没有推荐一下的等我有钱了鸡蛋仔，热狗，地道肠我都要点两份我不是爱情切不要哈哈聊聊天就好本来就是孩子嘘不怕你们笑我
欠的不多了，这月下来不欠公司的了哈哈林木朗赛特区内天推可上门接赔付的dd，两个人，可做客服、推广，正常离职，可开离职证明好困啊哈哈以前夜里打ag估计是你喝多了嫖娼票没了客气在哪里不用学习我我不管，我熬日批华子不好抽kkb完美还在企华园区吗？1累吗...666你想多了这里都是我小号明白了没兄嘚你吃的哪款上次行李箱大概10公斤
，80我怎么就不会用到处就有啊。。。菊花镶了砖吗卧槽谁有智栋物流的微信或者电报都跳IP换一个苹果7屏幕大概要多少钱哦悠悠却客闹，会所套餐来两套虽然我是很专一的人，活好就行回收二手老娘们，不锈钢盆换不会吧你在哪儿这个鬼地方实在太垃圾了别慌心情不好去死啊跟你无话可讲别气那里有好看的中国小妹推荐一下兄弟们一瞬间心情崩溃
了没看懂那有时间搞那个呀。她吧记录删了600你加一下我私人的是呀哈哈哈哈哈你这话说的，杀猪的哪些一个月几十万还一起约去舔逼剩下我一个人了我舔那个是甜的6000回国就赚不到了吗？还出来推个几把而不是像男人一样在哪里呀我要骚的其实西港好好呆着玩没有那么多事情的我以前经常去的几个地方当地人和我都很好怕断粮吗菊花早缝起来了
我就是美女.上次还说免费哪里有没有不睡觉吗我要点外卖谁有淘饱的微信谢谢不贵9字头3800包夜还没赚到五百万没有脸回去啊砍你给你们会所的女的都干的明天不能上班���哪里？名字也算嘛饱不饱满？哈哈想你了，睡不着，我的小宝贝为什么不可以生个混血我是老实人云南女人还是漂亮的哟200斤哈哈哈哈，笑死我了就一直在特区带着我来这么
久了2100小曾呢十个在吗？借二百是多无聊？我特么崩溃了骚瑞护照还没签证呢感情杀手？谁有洋马联系方式射出来血了是什么情况他睡觉了谁炒群谁知道哪里店里有毛豆炒肉啊宝贝还好及时倍销毁了海飞丝越洗越多感觉啥活动也没根本不是搓澡就是忽悠你打炮大佬帮我推一个卖手机卡的呗优质qq出售超强耐操价格美丽量大充足9位二太三太皇冠10位
二太三太皇冠还有大量业务号出售可以预计广告好友承接一切解封业务人脸扫码短信你值得拥有！！！还是说你是不带把儿的？哪个憨憨删我女人骗钱骗鸡巴可能你家物业不给力嗯不要会晒店吃你大爷日常签个到精神着一块非常拿捏我一个月就五千还TM要去包夜出去抓人了包夜滴滴安排甜蜜�‍♀他为什么这么色情我看看1我这也有要不要不是因为什
么麦克风给你讲出你的故事你是男的女的都是十五十六的什么果冻玛卡悍马糖狗曾多大有肉的私聊我哈哈哈推荐下你现在能来拿不有没有老板知道逍遥宫的联系方式或者妈咪呀问问你啊，今天白天都没看你冒泡可以吗说的好打条钢内裤你算个什么东西偷鸡巴都是拉皮条的各位有特区手机店的微信或者飞机吗？买两张卡用下泰国超市�路口查车，请各
位相互转告动脑子想想吧！黑户过来这边找个小姐嫁了吧WY的工具有木有卖车的本來就是12170好像马来西亚的还是哪的逮到给他电疗华源好我怎么了哈哈哈这个不存在那好啊@hzz9999聊啥她喜欢帅哥天呐伊利别怕，我保护好你感觉你不上班不在了不是要本人护照他肯定交了502是小气只要需要我的地方你帅！你不是这么害羞的人你去不去�欢迎加入
金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主有护照我安排那就好我马上偷渡过来还得偷度回去？现在如果要过来！可以安排到万象了！直接飞过来特区隔离2小时30008小时5000不给走我就睡地打滚赶紧祭天吧那我聊啥如果不是被坑哈哈哈哈妈的本地人有些会出特区�欢迎加入金三角
唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主我睡了多少一张那你说吧，到底要打几炮就好了群里有苹果签名商吗，请联系我接赔付，有无护照都可，茶水两千你又不付钱我他妈的不上了打算什么时候回家？那么好奇一起隔离感受哥哥怀里的温暖狗头们没人说话呢。你们干啥呢飘到人妖你也
算是很幸福的了给我推荐点好吃的可以吗我在特区的时候天天吃食堂嘛难道1月就严打1.4本人这两天要发工资了，有没有800块的妹子了解一下呢我自己进去为什么要了解被西双版纳抓的？我们那个县城的回去护照都剪了八点呀他们说的是真的吗卖电画咔，薇X，支扶宝签到我啥都没干啊我就玩坏了两个手机还赔了你已经凉了全部拜国内上牌了吗你有
枪因为发广告我怎么收钱我说的出做得到没有哈哈有好看的没有推荐一下卖了电话手表就成摆设了。磨砂是什么对什麼性质不一样带头冲锋@LYsoulmate桃总@aotemantutututututututu找你左前方那个拿一下就有了但是马上就被拉走只有进去了，才后悔当初要你吗我都说可以发下班了，工作号退了，这是我私人号杀你妈骗子工作室卖了去找小妹不香嘛
你来跟我混吧一个月四千难受什么尾巴给大曾甩一个啊特区外面医院阿不你已经不在了我请客安排你吗人事联系我妈的我擦我找保安借三人行免一单三人行免一单需要联系我安排会所中国妹子大曾我怎么绿帽子了？我一个月四千工资jimiwaimai778私聊我哪里知道不然回来你客户找不到你她在睡觉我去下个天刀玩找新疆人买东西，问价格不买会被砍贵
啊这边有优质的男人？你jiji都没有我口活指法也不错那我可以拥有你吗我是个纯爷们男的提不起兴趣来你是在特区吧众筹咱就算出的起钱人家也不会卖身啊收华为mate30pro5G9成新以上联系我有个梦想去找喜子我都拿了4个月的保底了�坦白就好了没微信�是的为啥要置顶语音我马上穿衣服这里面几个？你数不清楚？快了问的士有要房间的吗？jsba
by829012人呢公司赔付太高了1999收购一部11？老师可还行~奶子打肿我见过一个牛逼的，喝多去日批！射在逼里，还要去添逼，完事出来一只在吐，回到大黄楼刚下车又吐你在哪，不好意思1702不能在多了一出国内就知道你去干嘛了从来没去过不怕太子不是在西港嘛不堵车估计三分半企华第一后卫请求出战金三角没有土狗的话。明天请你们吃狗肉你
又发骚还能在哪原来你是女的受不了也没招阿哪来的什么60分钟看看情话在哪嘴上说有啥用小哥哥是拉皮条的啊小姐姐好不好看你你耳朵好使啊我花400块钱我又不玩快手几钱收有没有寄快递的联系方式我很喜欢那种一开始不肯，然后后面就给操爽的那种电子烟闻着香香的华源有会所？那在哪里哈哈哈哈恩，但不贵。最贵130你那还找得到吗我以为你
嫖娼去了还不如去吃汉堡我刚打开快手看的我曹2000多挺贵了那妹紋身是一個藝妓/�那金洗果果不过我没注意你啥样，90分钟的爱情比较适合我为什么我主要实名认证？正能量他们这样违法的……东北的怕这怕那。研究的好多啊小尾巴讷曾经堕落过黑暗这个奈比多各位大佬吃早餐了没有我们办公室的一朵花500三炮这他妈是30秒一炮吗。才能便宜成
这样好看吗阿巴巴？大曾睡吧……�组奥利给你还不睡等死啊是不是报备了就把你拉近隔离群了你妹双飞好可怜哈哈哈你下半夜那个M3等这个泰国妹理我螃蟹你拍个我看下嗯没开害连中两期的给我中3个亿一祝下了200有招人的吗？一名卧槽说要咋的�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天
在线发广告联系群主对啊，你还没睡我出不来？爱你吗我就是�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主周黑鸭？可能是真菜吧。哈哈给你来一份嘛细节不重要重要的是你发的所以代表的就是你水果微信推一个30万赔付谁接？嗯嗯提成多少@LYsoulmate哪里啊哦。。加什
么宝我同事很尴尬我老板他们都在…发出来。这首是我之前最喜欢的谁大长腿累就对了她出套子谁老艾特我我一个月四千工资特区招人，公司直招，接赔付，欢迎跳槽！！！！我要睡觉觉觉好的怎么会把这么可爱的小姐姐踢了呢不好玩老板娘在忙我还在敷面膜好的ne木棉诊所隔壁木棉酒店右侧老司机不夜城买菜不要了兄弟，换水还是不错的，时光驿
站也就20多个，流水？那我们刚好四个人放起哦哦冷清不会送毛啊？我上次来怎么啥都没有胃口真好哈。你是说梦里的花吗嗨，集美哼2个多月了这是个老女人。想女人想疯了是不是超级签您老请说签到你是小姐姐吗引流广告词改了哈哈请问有人在这边办理过驾驶证嘛为什么我QQ能正常使用？天涯何处无芳草何必单恋白班到此结束晚上继续好想做一个
美丽的指甲来我着再不行，直接把头伸进去白裤子个鬼爱会消失的实力第三方，需要的联系，可面谈，可押金担保便宜一点好实惠现在金三角的逼怎么不值钱了吗没疫情还好说你的小几把太小了不够吃没人接就退缩刚好我带头隔离费加核酸检测下来3200-4000哈哈还以为是豪志的可以一直来我知道哦特码27鸡掰我带把刀我也听不懂我啊我啊狗蛋我买了
。。。。。。。。我没权限空壳管理拉黑了哈哈，估计发的色图不是啊，你说你要把门面给樱桃了我以为你要走了呢非常好吃找你要营业执照呢一晚上不睡觉的嘿嘿熬夜对身体不好公司有需要养微信或者打电话的吗！本人找工作后面那两个小姐姐不错红迷7a我蒙蔽了你包红包了吗对象是对象老公是老公本公司多岗位直招，有意者私聊我哦我说你公司
是做什么的吗我没看懂你说的是什么。交友软件被。尤其豆腐脑，合适就干不是我不跑我尝试一下毕竟我是吵群的每天都有不同的老挝妹睡对的奶熊奶熊的睡情趣房不去打死好晚上发房间号位置我把失乐带着来晚上这红包不包咯?那1万你们结婚马上给哈哈哈哈死了啊？宝贝我给你凑点今晚？接接接，特区天推接跳槽，上门赔付害我才6岁樱桃你真的是
…滚这期做完就没有赔付了，我就可以自己卖自己了，去个安逸的公司。金沙国际会所.❤（阿東）越南妹纸又吗之前不是说好的没提成给吗？怎么现在又变了？是不是变得法的为难我情话！有人事嘛，来聊聊为什么我喜欢夜晚行动今晚最后冒个泡泡。你穿吗我其实不喜欢多人，哈哈6666每天还给我电脑，让我卖药这边有很多滴这。。。。。。平白无
故你说我要是说可以那我不比肾还虚吗����多少钱先发来看看宝贝，晚上约起晚安(☝｀˘ω˘)☝我想多赚一份钱出来聊天呀赶紧退群打游戏谁在部落签个到你用过吗消费不起鄙人要饿死在外面了当然@toudhuiguo这不是骂赵主席的大哥嘛你见到过？中国妹老子捅了病毒窝么，天天给我发求购二手摩托一辆安排约出来了请你吃一周火锅不好真恶我给你
家的温暖等太久了基佬我喜欢嘤嘤嘤为了能不能进门怎么了你这小王八蛋3炮来给你开个代理在国内混混吧塞林木个寄掰我先睡了你们聊我都还没睡我在财神俩人一起隔离你去蓝盾里面玩个几千万六合亏了4天饭钱必须罚大哥我错了开那么好的房明天我可以去性感小丝袜。。。群里有美女一起去吗47——35——34—37—46—27主六码07——10——11—
—17****防四码我带你8码老板要跳槽废了美龙湾很长，我老换号老妇女都跟他说的差不多了广告贴脸上了我不配吗害谈理想聊人生？有没有卖蓝七的咨询下老挝妹狡猾不？就等出来交给你绿色嘴会说怎么会卑微了你说个时间吧大曾关你屁事我帮你问问我晚上去你排队我可以打一个广告吗我听这个怎么怪怪的版本这能看见吗有柬埔寨兄弟嘛�只是没人
给我叫让你占了个便宜哈哈叫樱桃拿餐五分钟左右这么卑微？没毛病吧你太聪明了‘她老板是狗庄恩？原来什么你一个女的我一个男的讓群里知道还以为我把你咋地了你这粉丝不得把我撕碎了哈哈哈哈哈哈哈花总那个不在CXK??还有唐人街里面那个穿裤子更好照1米49呢吗的1月份的机票在4万左右草，还不如去带小姐哈哈哈搭白牛噢噢噢真的不会封打
死老挝卡不实名的啊.....酸辣的哈喽有没有洗车的没赚到钱不回去用心了。谁吐了老子7500啥都干其他公司也在陆续搬陈老表火的一笔我那时候走的路线这都被你发现了辣灶哥们明天去你那包夜玩一下再回去回国大概要准备多少钱我也要出去啊白熊呢用手机听问题是给够了钱也得从你们身上出啊想回去也能回去收u那是白送我有你放屁吧我t了，微信
群吗行刚刚在斗牛坐车就可以回国了带你的人没跟你说清楚哦开车接你们出去！不用偷出去你们看发的疫情通告了吗。你什么时候带我回家哎人家都走了爬山爬的啊不怕中毒可以！刚才听一个大佬说的只要国门不封你长得很帅吗回不去了？报备就是遣返�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时
全天在线发广告联系群主把你卖了当鸭子哈哈哈哈萝莉不适合你。以后不要萝莉了就为了赚点钱给男朋友买烟微胖才是极品。是啊？自己泡个试试看�有点老我在开房叫你家小妹帮你按按。。。几楼？？收现金ZFB换负盈利还可以对刷啊打折吗求钱没有非要当兵骗你不是人，昨天一个都送回去不到期啊八九千是正常价格说真的哈哈哈别了�隔离费两个
人一起啊四千搞定了特区外到国门一万……我感觉你们好像不需要我害飞回去的话跟我就没关系了我这里几个逼白天都是去找老挝妹五百块钱睡一整天我去江西特产烟极品金圣有需要的滴滴我包夜滴滴安排甜蜜有人想回家吗幸福呢你长什么样。要5W弄你这岛上除了这些怕鸡吧受累？不敢去经得起多大的诋毁就经得起多大的赞美文明你我他幸福靠大家
可乐的尾巴哈哈哈没有小姐姐好无聊没有对他刚才说要去有多少烟找我@TEIQUZHANAN这么忙@rxsmw18被你拆了还行吧我给你们打折出来吃樱桃这是你想来就来的地方？下次记得带礼物不管你是老板还是物业朋友的朋友我他妈的要有空啊哈哈�我是女的6个人多好你可真骚我在找大曾我叫凉薯哈哈哈红莞会所A牌原价998预约红牛898M牌原价1498预约红牛
1298Q牌原价1698预约红牛1498T牌原价1898预约红牛1698S牌原价2198预约红牛1998进门说找（红牛）就算预约莞式水墨90分钟不限次包夜按预约价×2（6小时）包夜老挝时间2*30以后开始（主打中国妹子）不报（红牛）买原单本店提供免费饮品，车接车送也可以送上门营业时间：老挝时间19点至凌晨06点VX：a2772800438app：2095873064飞机✈️群
：@jinsanjiao888888樱桃在这里我想提醒大家一句：当今社会互联网已不是法外之地，希望大家不要跟法律打擦边球，更加不要昧着良心去赚钱。如果你们非要干，联系我，谢谢！@Chenyu000111222领妹妹不有妹子晚上睡不着找我你不知道新人进群发言五句才能发表情包和图片吗我懂的鸡巴的，我以为你们都死了呢大曾带上钱我马上下班了一个十万
招几个一级代理一般人都做不了那个动作哪个啧所以都让开各位大佬群主这里樱桃你变了签到瞬间没有了灵魂广告贴你脸上了但是我钱用完了祝大家有个愉快的星期天��好风里雨里我去接你你特么哈哈哈？？？？？？？？？送飞机一堆10到20那你做梦会不会叫他张译我配合度极高以后还是不敢干这种事来吧不喜欢处男不想负责他嫖娼被抓了有赌场
的公共或者洗码的朋友吗全新技�术渠�道，精准用户引�流，转化高，真诚合作。我一直在等着你有没有小姐姐要做指甲和睫毛的呀有没有小哥哥要做面部清洁的呀去他妈踢就踢了慢！手机可以哈哈哈要死掉了勒一看就老挝妹不是拉皮条的吗走组队狗主管哪敢哔哔输了一百多个看过你请客是不是加我和我谈谈啊一直等着你哥哥,能满足我吗?想买个
苹果x以上的这里哪里有卖象棋1都为了挣钱兄弟收发验证码正常不就行了你要还是和我问价回不去的国，不如留下继续种菜媳妇你真可爱那在带带我咯哈哈哈冬天来了找个女朋友环境还可以吗？啥都喜欢3.5个点你在这里说国内。。。吐血钱在人家手上你还敢皮一下？怎么t啊我要你明白嘛老板有MODE私服嘛？人家买个车，顶咱们一辈子的？？？搞点
副业我看到有些外卖都发朋友圈可以送不用担心没事的，出去估计就可以有人免费带你回国的工行双密来跳付60！！！哈哈哈我三月干到7月知道我一天上班多少小时？？？结果把他摁住，草的哇哇叫有没有落地炮的多少算有业绩啊？入金治百病。不然晚上阿进又多几个客人了我要去干拉皮条塞尼木，发给他你是不是傻我母鸡啊哦终究还是开车了给了
多少好处费我去问问保洁1狗贼妹子你是长这样不到时候白熊组织下群主气死了也是什么雪糕我上个月月底过来的告诉他戴套大腿弯膝盖弯挺不错的你这个豺狗嚼的四国联查滴滴滴早上的消息不到晚上夏天哈我经常去博翔那异地一年差不多意思偶尔玩手游要是说那个真的对我胃口吧要钱也要试试不要跟我硬碰硬，我受的是伤你丢的是命我觉得找那些可
以蹦硬的以前打的老狠了是的一年吃不了5次伊利我也不知道，但是我就是想约你妈的热死人那么多国家都畅通无阻明年跟着大哥你混@xiaoyu6666666他骂你怼回去这种组长也有，擦咋了心情不太好没你文化高你就是敢不去冷死了我说好的，我说你喜欢我吗，她说喜欢，我说没钱，她就说妈妈病了你去了，早上起开就包夜钱出去了难道女的就没有丑的
在啊�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主怎么这么多皮条浪挂习惯了是哦，被逼无奈，就打电话给派出所了我说的是联系版纳的派出所，不是老挝警方我差100多万我都不急你急个锤子哦哦哦�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理
�有事找群主群主24小时全天在线发广告联系群主睡一起夜班�那赶上的时候不好大曾你不会出去包夜还想着工作吧脾气这么坏？你不是要跟你男朋友聊吗忘记了多少钱�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主被逮住了本来没有什么事还搞的我像坏人一样你在哪里打呢
赛里木�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主你手机里面保存了多少人了人家未成年哪有给你喝我想手臂上纹个简单的那算了，不和你去了出售P图软件，支持生成转账单，各大网银回单，以及各种短信到账信息！你是说报备出去还是车费？这个腿毛她在她在就老挝
时间下午4:50哦我想家了。
3万加个人删哈哈泰国还差不多搞的咱们好像都是坏人一样啥事老挝痛比较好很多哎主要是不想骗人什么的钱花着不安心���心软也是好事怎么会被抓省房费有些会中国话额真好吃了我两天工资法律哪条规定女孩子不能说那个了，群不是用来聊聊天难道是潜水的吗好的我知道了抱着老子玩电脑也才中病毒一次电梯坏了两次也跟我没关系啊我啥也没干
就去了两个月我见过。但是我boss好像让人抓了。总是打错在你楼上刚睡醒不是早安吗？把我的删了我的图呢？谁有唐人街老西安微信。推一下谢谢去死吧分手就诋毁你名誉我知道了送一份到缅甸小勐拉，现金支付你水果店在哪里啊你他妈的最坏的就是你樱桃是你能亲的嘛�理发店还有不靠谱的�大其力对快吼渣男要睡觉了请我吃吗小废物是真的废
物吗要是鸭子不黑怎么样kkb不吹牛逼了眯一会一会下班了。小藤条，水牢我来了不是这里的大家庭了嫖两个就算了看见了我真的知道这两个在哪里我便宜她两不用带是她们两带我出去晚上唐人街喝酒？唉刚才应该去豪志杀你妈大曾水好多哈哈哈金龙门口这边有宠物店吗金沙国际会馆90分钟不限次全套莞式服务带水床服务.找阿東预约.大家小心有我就
喜欢看热闹你这不行阿，要买点药吃@guanxi688888好的老板你都不验验货有车要卖的，私聊？去蓝盾找李哥他虎，我是他小姨谁有超市飞机不我洗完了巅峰2000以上才有资格大白马等你来土耳其种菜��  那你还问儿子乖我想偷渡029是啊你声音挺有磁性啊梦见被移出群昂我要删了号发一下，谢谢我可以熬死猫头鹰�你信不yes终于下班了香港卡可以注
册吗你消停干你活得了一查就凉准备出门了现在回去一个多少钱�签到搬砖啦想我了吗鬼子进村了什么连那家会所服务好他都知道我是楼下看大门的那还好买的吧什么口，今天嫖娼去了账单三年我们公司完全没有这职位我也缺钱你找我有事干嘛不是你发信息给我这个才1900，兄弟高矮胖瘦应该辣四川口味日常签个到都一个多月哈哈如果实力允许的话
卖军火多优惠上个月关了又不是一定又要你签半年偷渡过来的没有入境记录，咋去签证？冲冲冲儿子，吃饭啦我没带手机，你给我买一个呗有实力就拿户跑呗不能大声没有你退群吧带我去嫖在家Q看来是没有主管的希望了嫖娼能不能赊账身高一米七，腿长一米八这句话啥意思我辛苦半个月的消息晚安断电断网来跟着我摇起来��各位大佬招什么人哈哈
哈，包场要么就别出来我只知道给蛇头是8000快跑卧槽虚拟平台，走完全世界只需要10S这还用问？私微诚信经营群里禁止打广告啊你干嘛撤回我消息孙悟空哦我是说这群里又有点吓人具体没问我只知道他回去了不知道有多的没估计预约完了谁找我？马上给你罚款教育牛批不然我会去背包？哈哈哈你要是怕，你找渠道不就行看要的哇上人了新一轮开始
了掏了掏口袋的红双喜唐人国际酒店套房出一间，需要的滴滴！价格美丽！�不费肝你太胖了养不起十个萝莉九个富有女孩子玩拉拉吗��还开vpn哈哈哪有房子租单间就住十天半个月那种据说老板娘老漂亮了登下哪位大佬有迪升药房的微信呀，麻烦推下对呀阿高哥哥你要坚强表哥我听不了语音小萝莉都不跟我聊天==我的确说那种话不太好，主要是男
人之间的吐槽！！！我不知道你在本人身高170体重190，我长的丑，但是我很温柔不去带阿女人在哪里塞尼姆怎么就开始了mei听懂哈哈哈哈刚才唐在办公室看监控看你们小萝莉好的昨晚一档直接冲起上次你跟人家喝酒喝的最起雾的也是你老子喝多了，回你个信息@power89966说清楚尼玛的有没有回去的小姐姐一块隔离当你大姨妈来了我还可以给你递
姨妈巾估计是梦游·花400但是没时间过去我想过去，太晚了刚下头手机都亮着这个可以勾引我礼拜六你自己解决小全杀、看我·名字老板娘给我介绍霸道的小哥哥吧这种男人要不得你导个几把我不要你觉得我要我觉得�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主那个公司
的小米直接买新的就可以了对啊，住的缅甸啊ha啥我就说说卖鸭货是我的职责有没有要人的这就是冈本001原产的吧害腰疼往死里整哪个稳住，我们能赢叫我去啊我东南亚第一导师给你指导啊难道是昨天晚上樱桃又放你鸽子了不我没有你飞机也可以我说出来，你都说不出来好的华源了吗，这哪是海外，股票。我是一个炒群的，我射的快哈哈来我这里西
瓜霜润喉片�都敢这样招人了，还不敢保证赚钱别找我哈哈哈你在那我去蹲你最重要的是保护好自己的叼那你现在出来啊好的我日你妈也不要迟到包红包就溜？渣男是啊我不用排你不是白天睡觉吗我以为我已经很久了白班阿通知的草我还没说干啥呢喝酒啊哈哈哈哈别提他了有空我们小酌一杯？去拉皮条吧别推了说那么多把你的照片扔出来看看你看这
个只能在二楼徘徊本色就在唐人街大门口进去一点又没见过好想发我只想白嫖在广州就二三百我也有闪光就企华里面刚刚签到签到...................有你花钱啊那一天定位万象算了�卖了不给你护照？五百块钱一个卧槽比我这还贵你想笑死我最近vpn很不稳定你别整个男人啊鸡巴太大了，进去也不好进去给她口半天带带我今天的天气很冷，门外的
砖头依旧烫手，我先去搬砖了，你们玩我朋友说痒哦，退下吧去了七八次奶茶一杯，快乐起飞报个小到肯定安安静静的当个吃瓜群众你这样子，你给我找，不cao出血我付钱什么新货今晚怎么没看到那个猫呢���我需要是业务员自己的三方过认证蒙在鼓里嗯嗯，到时候再约，我这几天也忙，再调整方案曹尼玛你卖贵有点我感觉还有人要，卖350我感
觉没有人敢买17—18年做人事打了很久廣告你的名字太咸幼女群支持试看不白嫖我给你都是一个道理想在群里泡到妹子可能吗会所想听就听超级签@akmabc哦还好意思说有没有解决办法亲做大做强我来找你玩啊谢谢佛系�感觉你很亏的样子踢了这些发病毒的啊那我不敢买经验老练嫖完没钱给的话喊老板来接人找他要钱上的都是色粉。正常伟哥吃了是啪
啪的时候硬悍马糖是一直硬哈哈大量空闲爱过螺蛳粉外卖谁有叫这个大增的狗管理买阔以呀请我啥逛热死蓝带买不到那先把群主干吊吧晚上旧款河里很多本地人在捞6500留着买12不香么情话800包夜安排害8楼是什么博来着只能自己处理还好，我们点完还送套房子妈的我是那样子的人吗哈哈哈哈哈哈快则3个月满则10年潜伏期这边在跟别的公司讨论这还
不简单说这话下班下班厉害了这三四点这么还有老挝的在查hz给他看身份证也可以点数多少，你也知道快15号了啊，哈哈口味是真重呀你在马尼拉还是宿务是啊此时此刻此景，请允许我高歌一首找不到在回来做狗推这跟线人似的，不得请我来个至尊头牌SSS的你换头像了优质qq出售超强耐操价格美丽量大充足9位二太三太皇冠10位二太三太皇冠还有大
量业务号出售可以预计广告好友承接一切解封业务人脸扫码短信你值得拥有！！！这几天特区不安全吗》有没有什么消息？哪个手里有现金塞林木来一个啥时候啊大曾你是真的欠打明天给你升级寄不过来张的跟你妈的瘟神一样又不是做钓鱼的慌个毛你又不是白班哈哈谁有电影院那边手机店联系方式真上家不让你走我们也有办法老子腋毛都没有。。。
。。。店是在那里哦老板娘你们休息多久T8888果冻新鲜的果冻小宝你瞎几把乱发梅龙湾不知道那为了奶豆还是可以戒的管他什么正经一会要我继承她的家业岂不是凉凉7005多少钱一颗就让我将这份美好埋在心里白天打个广告各位老铁为晚上做准备要新鲜的刚摘下来的哪里有人物包装图片这些卖呀有啊在这里除了狗庄给的钱就是狗庄给的需要带妹子上
班没有你退群吧安排预约中.1158造谣还是你们拿捏的死死肯定刚出去一个小姐电梯都是香水味用手搞不动了你完犊子群主那跟我一起来做兼职啊明人不说暗话，我就要骗他2楼吃过了��卧槽我最近学会了举报飞机账号跟群能做到秒封西安那家面谁有微信打算以后在这边开�店，这也还不错？润滑剂？她给我做的呀我淦你嘴好嘛这就是仙丹西港或金
边屠宰场来个能说事的，有团队。诶？两个人？是的都是这样，只是要个态度溜了溜了啧啧啧狗曾这三Q调调到位了卖几把春药谁有樱花屋包厢联系方式大宝贝晚安他说你懂的多���了9开头我送花生米不知道。不嫖不赌不黄，我只想安静的当个保安不能推他老缅新村有个人房转租的？整的我也很被动淦多少64G的当着我的面，侵犯我？爱困快叫我美
女此时的我操是形容我惊讶，还有表达我内心愤怒的情感真尼玛口味重他们外块啊脑残，还怂。没进去前跟我吹牛B，进去了话都不敢说一直发抖叫今天新到的那几个出来上班谁有烧烤的微信来一个那个跟你说的好像是装没回家的想回家回家的又想出来我们有通道，直接用卡收或者直接走币迎着风接跳槽。无手续回国可安排你不带我去玩吗必须的水果
也会吃到里面有虫的我有我在金三角M3叫的电你妹滚滚滚不去了可以不哪里可以赊账？嫖娼回来了吗你哥谁啊凉有护照，出关之后呢？需要多少？老板，今晚又跑哪里正规上门按摩的老板出来一个上班呢谁有苹果11pormax原装充电器大家都是职业骗子晚上跟着你喝两口酒恐龙我也见过飞机带翅膀哈哈哈我到现在都不认识路我还没睡你不来我们都不睡
我就搞不懂1靠谱的来。我想吃这个粥几号啊这个我刚打完篮球回来高价收购手机卡，有的私聊我！吃着水果瞎逛嗯呢，好的要不要跳槽多少钱卖兄弟？我擦我的记得飞机有很多这样的频道，我找不着了晚上好连抢都没有盗我图都数不清几个人？我从来不承认我是萝莉天天喊我去打炮800还是黑T恤那个你加这个试试谁有zfb出啊�我现在就在吃�你跟
尼玛傻逼一样露天那叫上头不可能下头的那这样子操作什么链接？好玩你们这群王八蛋啊咋卖药的不敢卖假的继续干~花店微信推一下，谢谢�狗管理呢老子的小号qq前天换绑的这边的卡废话有啊不贵各位大神哪里狗头都给你掀飞真的吗我想请问，除了企华，金三角还有打篮球的地方吗要不要拿瓶五粮液给你打打垫底那只能提水果去医院看你了给我介
绍个女朋友，不要钱有没有人要资源的.前段时间，我朋友128g的拿去手机店卖2000我这边朋友有个11那个店的平常认真一点，需要冲么？今天，应该是半夜开进去的，晚上6点多在捞，把车捞上来了哈哈夏天是色彩斑斓会罚款的穷我说我我这听不了有没有云南口味的不是鸡掰这句话应该我问你我自己解决吧那应该在哪里问好一点哈哈晚点看看能不能
约到妹子呀�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主这样就废了我不我就要现在帮我安排吧�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主本群还没开启抽奖打飞机没老婆比较好全部挂壁你是
大可爱。�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主168是幸运的阿怎么了大姐你店在唐人街哪里……签到好看不哈哈哈哈几栋几楼？压几付几？要罚款我说的是实话你不是一个月就给我一百生活费吗给我置顶后悔必须的我睡着觉哈哈做完性生活和谐吗卧槽只要几把大只
进人身体不进入生活���又关你的事一阵一阵的你看她id谁老艾特我…收张国内实名手机卡只要老卡，有没有卖的你晚上跟我讲没办法本来就丑了还低调那岂不是一无是处在某些时候我还可以叫爸爸好的呀等一下哈280在引起你的注意害奔放了哈老板娘三两个橘子送吗……近来可好总盘2.5下午那会都快炸了樱桃睡了吗卧槽要不然叫老板娘送来一个
女孩子吃宵夜黑名单？我不念旧的我傻才把照片发这种地方傻逼我眼睛越来越差了，现在到了几乎看不到的地步了，我想在失明前看看身边的女孩子穿JK，不想在网络上看了……切钻石哈哈，我要包夜去了就是那个叫宝贝的据说女的吃了伟哥性欲强的很牛逼也就罚款调查隔离只要出了磨丁口岸，就别怕不过他那个渠道被抓了我见过一个越南女的，偷
渡2个国家。被抓了判刑2年爬山七小时真的假的谁大曾你就是想骗我发手照和腿照照片发出来给他看看西紧啊巴好嘛今晚不吃五排那我怕要在床上躺上半个月才能下床我草需要租房的来两房一厅一厨一卫因为没见过牛批30如狼卧槽你这么自恋啊一直打不出鸡巴会肿起来的顺便帮我打一下打飞机这么多人在这里，我会含羞的搜索就可以完事儿。女的做
鸡，男的做鸭咯我请你才19岁必须的阿网差的一批关键没有啊怎么跟昨天不一样卧槽....忙完了讲道理为什么�没有乳胶到底是怎么算晚上安排哈哈哈哈飞机杯充气娃娃价格已经订购了低清洗媒体信息流可上qpwz国内到版纳16特区一条202条40挺急的晚上怎么安排兄弟，辈子兄弟大家好不知道白嫖到这里了，可以啊兄弟特么都成管理了加一公司直招女
财务，有需要的可以私聊金山角刘德华。给你请安了喜欢你这大钢筋你猜年轻好啊羡慕特区啥时候解封啊，谁知道说谎园这不让出去啊只是知道是个要通道？不是年轻人不讲武德，接化发，耗治尾汁姐姐我你不是想吓别人，是你自己怕还行马上又到了干饭时间了没有不去日常签个到租微信吵好骚哦动画的啊在国内一万还能豪横一下楼下就有开房的跑
那么远干嘛不是所有牛奶都叫特仑苏又要杀猪了你们喝奶茶吗你说的我慌了出小群弹窗股民fen大区股转彩可到场地手工代拉小群一手报表需要来聊我约过一个博翔的妹子太狗了然后现在早泄吗我也不懂谁有这个的种子？原来如此那就太远了现在还吃早餐你他妈的泰和在哪啊怎么说心中有曲自然嗨200斤狗曾，你居然为了一个大吊M妹这么说我们被生活
所迫亏不了小姐姐签到别再打了没睡好在哪里按摩哪里有萝莉就叫赵云都有打必须打露天酒吧咋了这是蔫了吗手真贱还在转啊转谢谢爱你大龙是谁嗯，我觉得可以谁知道这女的哪里的咋样现在恐怕不行吧我记得黑妈妈有一个服务员长的还不错有没有发仔茶餐厅的微信啊你在抢手也有人了啊你是在哪栋我来救你你要哪一家的我天天上忙得很，没空吵群
好玩吗所以真的是玩的时候别人对你是真的，你真的时候别人就想玩?���没有啊能6200苹果11promax256群主晚上叫两个妹子出来喝喝酒啊这个美女好乖啊晚上怎么联系哈哈，我不賣檳榔我们不再一个地方吗?在国内开始你的表演你太丑了不约。音响师有人喜欢啊onlyone心诚则灵等等私发我这怎么看着象男的嘞睡觉的事情就交给我吧。❤️��哦豁
不好听了我也无奈啊以前没怎么胖，酒窝还特别大计生协会你不是说我男的吗......哈哈哈哈贼喊捉贼各位大佬们哈哈玻璃心了哈哈哈跟她聊不吵起来才怪不要不要男男吗这个我不会呢放屁你@我干嘛，你要死啊我一个人自言自语嘛真的来来来这个号没动图本来就不多又是九州一看就是早泄阳痿的人了。鉴定完毕然后把你丢给狗庄狗庄在给他2W接着奏
乐，接着舞�辣上瘾的微信谁有今天没得空你给我科普科普也行有人要吗的偷渡第一人听说非常热闹有什么美食推荐嘛这房子不是淮阳楼的吗yes没钱玩什么看他的操作频率应该封的快啊他们都说罚款不拘留皇冠50个有没有男的没几个搞这个好位置我也想试试，但不大好意思进去还在联系求求你们做个人，敢不敢加个好友坑里是哪个地区的那个不是不
得劲嘛不好喝不兑饮料喝的吗我还以为像我一样喝超了18楼的？我也买个好了有没有三人行一人免单的谁有四川豆花饭的联系方式给下谢谢大佬我也想尝爱情的苦有微信刮了是的没错.给几个好喝的奶茶店微信，谁有。谢谢我跟你一样，兄弟怕骗只要有喜欢的，不差小哥哥没爱了，睡觉睡觉等下他们开完会，你要来再来凡事不要慌，手机拿出来哎那你
什么时候过来的？那么爱上网的吗关你有懒觉关系聊了个按摩妹妹，今晚要用餐了钻石的我现在过去奶茶那你按摩800爱不会消失的云平台收不到验证马都不用申请有缺代收渠道的帅哥美女吗医院？我都看到正脸了咋滚吧你我会的口味真重真好你猜想吃早餐了都捏疼大众了一个盘的证据300斤送外卖去了@alang112233我弱不禁风熊宝宝！！！！晚上好
♥️♥️♥️哈哈哈太难了大黄楼楼下不然？信不信我投诉你带你去修车10万才拿1万多“啤酒饮料矿泉水、花生瓜子八宝粥”太子说实话垃圾又来这一句他敢管你聊天，草能不能嗨起来先来一百个缺男人吗本来下月工资充值看看我不跟你们吹阿约个饭吃吧嘻嘻☺️☺️☺️都有点控制不住自己不是，是要中国号好的笑死我了羡慕你们那里还有冬天你知
道我在想你吗你也没干了？你跟前我不配有支付宝境外码的联系我有病才能去药店看吗又熬過一個夜晚还八号来。。是个男的狗群主管理啊不了回中国也不错群里有没有从缅甸到特区的路子？寶貝們二手老娘们有什么好收的两百一次不是我不是同性恋需要啥子位置掉毛留下一只手的话可以的我之前一直都是绿钻听接受预约中可以可以��这tm1月份我
去青岛玩，然后特地跑威海去见一个狗推跟你说，我今年看透了，狗推这么个东西500？号耐干那个卖号的教的都30多了群主把我变群名了^_^赚点嫖娼钱为什么上班的时候要一起叫口号不会你看我都没有看到钱太违心了，也不知道这因果循环来了是什么样的我要认真工作不能再堕落了有人回家吗当天出发当天抵达不需要任何证件正规渠道走国门不需
要拘留这你都知道？你开心就好妹纸吃了我可以玩给她改个网易云阿姨我没动可乐的头像我色诱他干嘛现在这里就留了可乐在这里了你们谁包夜的时候，叫我去舔，我在睡一下3我需要这样的女人。你咋说的我啥都没有你的不能信几年没中过谁有卖小龙虾的微信这个是新妹纸？你走吧没事，爱你tui1下个月必中全部都用不了了你要吗？不用三百可以给
你！要多少？这个图是不是p的啊我们可以白天去这个地方少嫖点，身体透支了越南的好�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主对刷负盈利啊你去吧特区内有没有对刷负盈利的公司我喜欢前凸后翘切，洋马没意思要回国了？一分钱一分货给他查我做正当职业你是大佬
打游戏去了可与可与害被搞了自首变遣送八点到八点啊信哥，砸，砸中了就有钱去买生日礼物了我一天工资才两百好好没毛病好的要什么？微信号商看到私我微信号商看到私我微信号商看到私我.现在不去赛里木鸡吧去的时候叫上我有事找我你要多少�开业就看你打广告了他妈的只要真的年龄这么小那估计要被饿死了世界已经大乱、到底是人性的扭曲
还是道德的沦丧还是你想的周到
企划对面哦给你说了你想破脑子都搞不清楚正常不过你喊价位高的啊网易云？弯的我喜欢自己去。别急不用在乎小尾巴好这个就不知道了我已经8个月没出过特区了。换一个头像a睡眠这么差吗渣男？多少钱一斤樱桃我操，九妹你把你p那么漂亮？老板走了我的天下了开车开起来各位老婆好，谁有开视频的软件吗？不带你可以但是钱必须要带耳机单卖不
就是不能惯着咋拍呀对面我擦华兴接人不我们公司有几十人在柬埔寨那边了伊拉克兄弟们护照老板那在哪里你有？今天爆满一看工资2500���担保有什么用哈哈我之前做两融的公司转币给你，你自己卖出去没发现9楼有夜班啊？黑暗杀踢了又是上夜班的一天还费肝�我是个女的1开车使我快乐出一部手机带微信是一整个盘好无聊呀哈哈哈来接我了吗C
anIseeyourreproductiveorgans别喝了做个人不认识什么东西真的是什么都推给大姨妈拉进来安我也想找一个在我觉得嫖娼的人都好有钱要天然无毛的治安局真是�西港欢迎你狠人错过了错过了通水道群里有人有手机店的微信吗不知道☘️☘️☘️当管理员第一次删消息到底是谁走漏了风声卧槽哈哈哈卖卡特区有24小时理发店吗？回去就被剪了已经注销
了被后宫佳丽三千怎么了宝我喜欢男的你还有老公啊，羡慕这事儿我去帮你办了牛逼看来都是年轻小伙���再发一遍你们现在多少业绩了，天天看你们飞机上吹牛逼做什么哈哈哈哈要啥飞机那种不甘的笑荷塘月色按摩手法行吗？唧唧歪歪i爬到一半，花生米就来了�没看到阿报了多少个表给我发轮发功了等着猝死吧大曾金三角哪个机场可以到要啊她
鸡巴给你打断夜班仔，周五不喝酒，明天变垃圾要不要体验一下添屁眼多少起送不需要自首的那种泰国超市外送吗你没去过韩国？擦白面你在哪里，刷多少？我爱吃爬山仔宵夜带我一个啊蚌塞，不是加塞抽烟身体好以前上班习惯了我什么时候要我也要多少钱周末嫖娼我请假了不要在该慢的时候很快�重庆柠檬酸辣粉又没汉城的？你见过啊我怀疑你是
个酒店前台能赊账嘛？我问你，来，晋江嫁妆都是怎么配的.....我的后缀谁改了我的后缀呢你下啊咦来泡茶哦！动不动就爆粗口好吃你请客吗喊起来你之前西港哪里啊聊天就好角度问题大佬们都是什么渠道搞得小摩托睡醒了？你爱爱爱，爱你妈个逼玩三炮花500不吃这个花里胡哨的奥里给哈哈国内严重了，我也是听说的，不知道真假奥利给哦这是什
么删了你长的帅他下次不要再问了不去了不要慌还好吧在买�上班了傻？可能是的第一次喷人有点紧张我房间就一张大床可以睡三四个人嗯最近身体不好不能熬夜接多高几点了47——35——34—37—46—27主六码07——10——11——17****防四码不是彩票的大胆砸下十万红包收了嘛人穷家贫，还他妈的丑胡辣汤牛肉汤豆腐汤你有什么几把心事乐总出
去玩不带我？今天没出过可以收账在哪里啊都是换汤不换药特区外有没有茶他这个头像不好看还帅鸡毛等下晚点出去吧不要那么直接好吧没事可以把我们当成丑女狗曾@LYsoulmate分手吧好吧张宇的歌好听没事的蔚蓝出钱吗不爱卿平身沉迷包夜无法自拔卧槽吃烤肉喝啤酒你不送？群主不吃了可以喝猛牛这不是一样的吗，，笑死了伊利那你是啥仔为啥不
能照，可恶烟都抽两包了你在发你就给你踢了我是男的还这样？那我也认准备偷老板50万跑路的知道了，没有去过，不知道好不好】发个语音，照片但是不是你@laowang666888�贪便宜，中标也快。飞机里面那有什么女人，除了狗人事，就没有一个真女人。早上好嗯嗯，我上次点的那个很好，听话，我叫他怎么说他就怎么做真的吗？直接抢懂什么阿
拉棒要在迪拜的话慢慢走吧晚上漂亮的都在哦在国内用踢了每天加班半个小时打底坚持一星期这么狠啊你看我可以不回收二手老娘们，不锈钢盆换还是中国妹子好一些赌场的老板帅哥要是赌输了不得我养前几天还看见了感觉不对就往嘴里怼呗你这个逼怎么也在这不然你没办法坐飞机的擦我得20个月了估计哈哈遣送回国罚款多少啊？你是男的我就想接
点特区回国单子赚点外快国内快要混不下去了他居然说你要死了金三角第一小几把应该能遇上一个人家摸就摸了你们也摸不到只能看看你刚才摸了？说实话，这个脱光了鸡儿都不会翘你叫妹纸陪我喝酒唐人国际580的房间没华源好@aa2020b我单买的27快点去买买4925262701我一个数买10块钱不中砍死你群主帅吗录上来让大家伙瞅瞅早打工人哪里有卖那
个泰国老虎贴的？小场面我知道了你们去私聊打K人，打K魂，打K也要大铁盆@TEIQUZHANAN出来吧分个几把贵得很这不是没见过本人嘛红包呢上午聊这个�咋地给女朋友美颜三大磨皮美白收敛怕是有误会不管咋样都要学我？我请假了啊我实在受不了你们这种到处蹭的我是男的头发肯定不长@TEIQUZHANAN等我这一波结束有的再见她男朋友太多了要不然一
天又白干这么早王祖蓝欺负大曾这种没对象的我现在抖音都刷不起来了花总才500？我励志要垄断金三角的qq社交业务?哈哈小姐姐带我做狗腿吗没。肯定是自带房间我都开始上班了爽一晚就好了爽一下也可以了有几个女的卖肉？卖qq?我2天一张卡的哈哈哈十个现金谁要，四个点，没有没有哈哈哈哈哈木瓜上次好像看到一眼哥哥们需要美女联系老弟给
你安排三人行免一单中国妹会所服务啥�没有不良嗜好我回国泡泡妹子不好吗这边你多大对啊有华源18楼的联系方式吗可以女管理？有么？？我们公司业绩一半是我自己一个人干出来的老挝女人多这真的是事实这里的性比较开放没得了精料一下要什么钱啊万一基霸掏出来比你还大呢？赔付2w多有什么办法解决吗不可能再说一遍带休息？有人拼一下下
午茶嘛�你这在教人@asd38567私聊ta我还是叫群主改回来吧我们一起来的其他我都没去过为什么就叫宝贝了？大那你快点我买尼玛六合签到好多人叫宝贝这叫你走吧，她不要你这也叫靓仔捡到换给我酬谢两千也可以啊我有你很多办法！没怂过给个微信呗什么我t了红桃你这个副业不行啊大曾很疑惑@TEIQUZHANAN都懂再说了其他群我也在你去吧操再也
不买了兔总有一次赚贱狗鑫景国际在哪里看看你发的什么几把我在国内首付了2套房子房贷都还没还完回不了就算死在这边有金三角人事吗签到不是说没实名制的卡号要停用吗你隔离都没有钱打赌吧我收拾一下�就怕报备不了翔派小馆昨天给你发视频咋啦对那我是谁误会你了！嘿嘿我同事回国发微信交钱800不是，750吗@akmabc偶遇大傻屌再加500Q排
包夜没事再回来。我等着你只能玩得起2500他要拉皮条生意太好了吗无视我了？摔坏去就不会嘟的一下了说真的我感觉福建人真的很会做对，录笔录是最关键自首回国会留案底吗？他俩点上班不找了打工是人上人不信我�����对啊在老挝就是拿钱办事是不是傻，宝清醒点睡着了吗封我谁给的小尾巴我想问你明天问我要的谈哇，这边女的挣钱多容
易沃日这个这么好看这么嫩呀私聊直接小草丛好的，只要你需要我就送过去有没有亚马逊牛排店的微信看见了可乐在吗舔了也没机会舔到最后一无所有我原来住一栋现在搬二栋了。？别在新加坡不是白熊弄的管理没权限他只会吹牛逼搞错了鸡毛wow。这么美哪里来那么多可乐我舍不得你这种情况怎么解决？不怕不怕压我手上呢我的套房没劲。。。。女
的我有十几个小时的路程吧我还在床上躺着你就上班了喝酒唱歌滴滴，卡座KTV包房都有那个家隆超市旁边还是个憨批狗主管的姘头，都要去草了她�不为别的，就为了想每天睡醒就能看到你招个财务，做过支付通道的�麻痹特区待不下去了这下想起来了出中国手机卡，私人用卡可以接打电话，特区现货，仅此一张，现在过金三角费用多少钱有赔付和
没赔付的不是一个佳对离职提前半个月心累了女生是什么东西我都不知道群友都有优惠特区内菜市场有卖大闸蟹吗谁有那个狗推根据地的群啊宝我也想要请原谅我不厚道的问出来了有好不容易弄个兼职没开始就黄了这是没看上啊稳定资金代收你踢我？1点以前呀做爱做久了那床上谈黄谈黄的是什么都有我管理不在了呀带我你在国内快去吧要约伊利的赶
紧约现在想想算了我去唐人街每次都是走路怎么了呀我送你一瓶安慕希不多无所畏惧如果你工作不开心，告诉我，我带钱上门接你买槟榔扫码蒙对了就狠狠超群也比有些人爱装的很别装的很文明收到中了换别墅，嫩模会所喝茶老板联系老弟安排三人行免一单谁有沙巴的上家说一声然后你也害怕也不知道咋办然后找了个机会想办法跑了没太多，最近两
个月十多万千万不要说金三角！敏感你是偷渡出去吗？公司报备进特区的这是别人是谁�全剧终夜班妞签到你连做梦都在想着工作了金沙国际会所.新茶到了.（阿东）靓仔是谁打你狗熊�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主相信我自己哈哈也不知道5楼那个妹子还在
不在我不熟悉啊你和可乐分手了吗�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主她是从国内和我一起来的可以快回来吧我想你了劳资手咋不好看啊渣男你还拉皮条去唐人街很多适合你的给那就快去真好招聘精聊高手，无需聊感情，10天必出业绩，做互联网项目(包括小白公
司有专业培训)，接受合理赔付，提成20到40个点，上班时间：早上9:00-晚上21:00中午和下午吃饭有休息时间联系：@bingge888立刻给男朋友发了个520�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�有事找群主群主24小时全天在线发广告联系群主这里接人那么多，我看最后都是我们几个送。隔壁有一个首充120万这几天
回去人少人才我朋友回去支付宝一两百万流水被查了就说赌场赢得叫赵主席送你走外卖群给我改一个如果是杀猪就不止是一年了有女的？好的给你加生菜包五串五花肉吃烧烤终于回来了包夜多少卧槽南昌话？南昌话谁能给我推荐点什么吃来别吓到bier嗯了可乐？柬的那个可乐吗？我不配为啥群里不让打广告勒，这样不是更活跃一点吗谁给你们的小尾
巴微信号商看到私我微信号商看到私我微信号商看到私我谢了有多少要多少果绿色6500你这个是什么手机啊嗯嗯@Mantovin一起好的晚上不是带我找老板娘吗他就是姑娘可乐在线接单他妈的、@LYsoulmate给你准备点心卧槽，金三角去不了没有上亿我直接通报地方军还有手机银行可乐放没放啊金三角人事还有不？我女朋友也叫糖糖周六晚上约好了，你
不来老王准备跟我兜风在养生哈哈哈我办公室30多个人也是满背的这种壁话都说得出口宝贝配握草不陌生小鸡吧要死我之前翻空翻的时候遭过一次怎么打杂？晚上好啊我这里也有什么？你他妈放我几次鸽子了？小姐姐都不跟我睡觉的我是打杂的哈？哈哈哈哈哈我第一次看到这么好听的小姐姐这么牛逼。。。给我来一份，送汉城来在哈哈海鲜城那边傻
逼叫彭于晏TM的你还包夜？水多的年轻女孩请，往后站银海一号公寓42.8的6500一月交一压一26的4500交一压一合同半年一签有意私聊好嘞有绑卡嘛不是伊利做好计划河边那个不过逼很紧的所以要提前预约每个周末都可以出去啊送哪里老板想你了呗正规的要多少钱呢？不然人家以为我也是呢1来啦来啦做好事不要留名银坑镇上的？SQS以上的看得过去
喆衍,[28.10.2018:39]群里有没做美发的�没有有头有脸的大哥常年都在华源和泰和住小雨是几栋的？喝的我肠胃炎犯了���你还是算了吧三秒男夜班仔要开房找我真的还是假的有没有大老板组团嫖娼的对周末咋们去酒吧喝说不定人家钱多把车推下去的难道你们有钱人不是这么玩吗，我又没去过，都是听别人说的对呀血都打出来了如果太远我都不
知道怎么坐飞机，我失信人买不了机票哪一家的披萨比较好吃的我就好奇问问而已你把我拉黑干嘛有联系方式？我是一个资深的厨子这不是黄色有多骚小姐姐还能要啥好吧哈哈哈哈水都干了不要脸。。。。宵禁不是解了吗你去约给妹子哈哈哈？？？？找群主南伞那边边防派出所太黑了老客户多我算1700他们说的是偷渡50那都在规划中我都不怕，你怕
个懒叫想当那不叫我们你应该送点给我吧介绍个长期富婆，给一条华子特区有人会补胎吗我知道有没有烧烤的微信当狗推吗？这么多钱多少人分你知道吗？不知道还有2天把业绩搞起来大家早上好好看确实像我不知道该相信那个正常还是自己翻翻吧别捣乱将就吧，500就是包夜发个朋友圈，就好了签到�会不会看算了吧对啊厉害厉害，平均每秒钟花费1
块钱的男人看来你比我还不懂偷渡来的，想走正规报备回国接受隔离玩玩有没有好吃的外卖你是男的？你这样会没朋友的开了你就知道了渣男说的那你凉了生肖有猫吗微信又特么封控了。群里有再找渠。道引。流的吗��哈哈这踏马就是梅龙湾我还以为渣男是想你想的睡不着�欢迎加入金三角唐人街百事通本群禁止:广告色情暴力违者直接封禁处理�
有事找群主群主24小时全天在线发广告联系群主我在我支持你们我还不知道怎么说47——35——34—37—46—27主六码07——10——11——17****防四码是价格不合适还是没有满足你的他不配空闲预约中.她来我们就玩她是的真你妹哦？忙完了老板娘我忘了我有啊来活了出来接客宝晚上好直接联系同事公司就好了我也想看@LYsoulmate喜欢你是觉得我
不该赚你这一千五吗？？？我有你就够了这都算少的了好勒可以我们的文明群没有你走吧做生意了300再见渣男在侮辱你的智商那么把今天的杨枝甘露安排上吧特区有卖蛋糕的么我特么今天生日给老子文明一点不困挽留回来了嫂子还能介绍?您想多了不知道你看他没精神了不信他们我还zhuma发翠花照片里翠花明天晚上跑步去吗花总乐总处个对象，来个
妹子，11马上就还要降价百度也不能乱百度难道都是骗人的？19年不敢看还做内保坑不坑加上我有六个月了叫什么嗨话岛上梁朝伟来了昨天落实方案罚款算在里面了然后隔离来个小姐姐都要被你们吓跑我没对象谁想大众我删的所以就不是一家啊菠菜不会是搞基吧，虾仁签到姐姐不要聊骚了我敢什么东西特别是我大早上鸡儿就给我看的梆硬可以啊，给8
总管下退群是吧不知道我也不会凶了你小宝贝爱你做狗推不容易啊是少之又少了小可爱你看着买吧666群里面的人，天天上班都聊飞机小奶狗都是自己射了不管女人哎哈哈天天在这里做白日梦不许糊闹睡觉外卖能叫黄瓜�....？？？？噢shiZ`
	m := make(map[string]int, 200)
	for _, word := range wordLis {
		n := strings.Count(s, word)
		if n > 0 {
			m[word] = n
		}
	}
	fmt.Printf("%#v", m)
	fmt.Println(len(m), len(s))
}
