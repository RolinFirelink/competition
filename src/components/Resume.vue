<template>
  <div class="registration-container">
    <div class="header">
      <h1>1月31日杯赛报名情况</h1>
    </div>
    <div class="announcement-wrapper">
      <div class="announcement-title">📢 赛事公告</div>
      <div class="announcement-content">
        <div class="announcement-item">
          <span class="item-label">开赛时间：</span>
          <span class="item-value">1月31日下午两点</span>
        </div>
        <div class="announcement-item">
          <span class="item-label">报名方式：</span>
          <span class="item-value">请加QQ群 <strong>743749814</strong></span>
        </div>
        <div class="announcement-item">
          <span class="item-label">保证金说明：</span>
          <span class="item-value">报名需缴纳10元保证金。正常参赛将原路退回；如鸽子，保证金将作为奖励给到赛事冠军。</span>
        </div>
        <div class="announcement-item">
          <span class="item-label">直播信息：</span>
          <span class="item-value">赛事将在"每日一灯"B站和抖音双平台同步直播</span>
        </div>
        <div class="announcement-item">
          <span class="item-label">特别鸣谢：</span>
          <span class="item-value">感谢<strong>盖世小鸡</strong>对本次赛事的赞助</span>
        </div>
        <div class="announcement-note">
          💡 本次杯赛参赛人员和保证金去向都会实时在本网站更新
        </div>
      </div>
    </div>
    <div class="round1-wrapper">
      <div class="round1-title">📋 第一轮赛事名单</div>
      <div class="round1-content">
        <table class="round1-table">
          <thead>
            <tr>
              <th>序号</th>
              <th>昵称A</th>
              <th>昵称B</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="match in round1Matches" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="rules-wrapper">
      <div class="rules-title">📋 赛事规则</div>
      <div class="rules-content">
        <div class="rules-section">
          <h3 class="section-title">一、比赛概况</h3>
          <div class="section-content">
            <p><strong>总人数：</strong>68人</p>
            <p><strong>种子选手（4人）：</strong>罐罐、AOZaki、学外语好啊外语得学、SiegeCascade</p>
            <p class="highlight-text">→ 直接进入八强，无需参加淘汰赛。</p>
            <p>剩余64人通过淘汰赛决出另外4个八强名额。</p>
          </div>
        </div>

        <div class="rules-section">
          <h3 class="section-title">二、赛制说明</h3>
          <div class="section-content">
            <div class="subsection">
              <h4 class="subsection-title">淘汰赛阶段（64进4）</h4>
              <ul class="rules-list">
                <li>每轮按当前段位重新匹配（强对强，弱对弱）</li>
                <li>采用分组对战，8人同时进行</li>
                <li>抢2胜利（先赢2局者晋级）</li>
                <li>每轮必须使用同一角色，不可更换</li>
              </ul>
            </div>
            <div class="subsection">
              <h4 class="subsection-title">八强及之后</h4>
              <ul class="rules-list">
                <li>4名晋级者随机与4名种子选手匹配</li>
                <li><strong>八强：</strong>抢3胜利</li>
                <li><strong>四强：</strong>抢3胜利</li>
                <li><strong>决赛：</strong>抢5胜利</li>
                <li>八强起每局败者可更换角色</li>
              </ul>
            </div>
          </div>
        </div>

        <div class="rules-section">
          <h3 class="section-title">三、重要规则</h3>
          <div class="section-content">
            <ul class="rules-list">
              <li><strong>签到与迟到：</strong>每场开始前会在群内通知，选手有3分钟准备时间。超时无理由视为弃赛，对手晋级，保证金计入冠军奖池。特殊情况可延至5分钟。</li>
              <li><strong>网络与掉线：</strong>选手需确保网络稳定，建议使用加速器。对局中掉线方判负。</li>
              <li><strong>角色限制：</strong>淘汰赛阶段不可换角色；八强起可更换。</li>
              <li><strong>结果反馈：</strong>每场比赛后需立即在群内上报结果，以便安排后续赛程。</li>
            </ul>
          </div>
        </div>

        <div class="rules-section">
          <h3 class="section-title">四、奖金与奖励</h3>
          <div class="section-content">
            <ul class="rules-list">
              <li><strong>冠军：</strong>50元现金 + 小孩签名版启明星手柄</li>
              <li><strong>幸运奖（抽奖）：</strong>除冠军外，所有参赛者可抽中签名手柄（比赛结束后抽取）</li>
              <li><strong>保证金退还：</strong>赛后3天内发收款码退还，逾期未发则保证金计入奖金池</li>
              <li>若奖金池有余，将通过投票的最终方式分配</li>
            </ul>
          </div>
        </div>

        <div class="rules-section">
          <h3 class="section-title">五、注意事项</h3>
          <div class="section-content">
            <ul class="rules-list">
              <li>不设亚军赛</li>
              <li>比赛结束后请关注群内通知，及时提交收款码</li>
              <li>如有疑问，请在赛前提出</li>
            </ul>
            <p class="rules-wish">祝各位比赛顺利！</p>
          </div>
        </div>
      </div>
    </div>
    <div class="stats-wrapper">
      <div class="stats-content">
        <span class="stats-label">参赛人数：</span>
        <span class="stats-value">{{ participants.length }}</span>
        <span class="stats-unit">人</span>
      </div>
    </div>
    <div class="table-wrapper">
      <table class="registration-table">
        <thead>
          <tr>
            <th>街霸昵称</th>
            <th>街霸ID</th>
            <th>保证金去向</th>
            <th>段位</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(participant, index) in reversedParticipants" :key="index">
            <td>{{ participant.nickname }}</td>
            <td>{{ participant.id }}</td>
            <td>{{ participant.deposit }}</td>
            <td>{{ participant.rank }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Resume',
  data() {
    return {
      participants: [
        {
          nickname: 'linrichu',
          id: '3062933800',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '我家猫会流氓后空翻夹头',
          id: '2783676023',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'M别打了别打了M',
          id: '3272404266',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '若有聲聞者盡皆斬落斷魂於此',
          id: '3576414025',
          deposit: '已免费',
          rank: '负分大师'
        },
        {
          nickname: '红色流浪星',
          id: '3498125113',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '一等情事',
          id: '1739018635',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: '两仪未来',
          id: '3876379621',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'AOZaki',
          id: '3431187004',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: 'Riven',
          id: '2696437313',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '左眼下起雨',
          id: '4226859931',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '我菜我不难受吗',
          id: '3766826973',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '拉面馒头',
          id: '3753462047',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: '历战街霸金狮子',
          id: '2662803340',
          deposit: '已缴纳',
          rank: '新手'
        },
        {
          nickname: 'chilleroo',
          id: '1897981597',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '鱼见千寻',
          id: '2177871520',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '也许不是不开心',
          id: '3791128655',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'Farid',
          id: '4282409574',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: 'AzS_绕远路 才是最短的捷径',
          id: '1315717156',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '我是一个超级大狗腿子',
          id: '1314720339',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '胆小的阿烙',
          id: '3538353124',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: '喜欢罐罐',
          id: '1324496630',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: '意墨要起飞',
          id: '2171623952',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: 'SiegeCascade',
          id: '1083334838',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: '长征永弓使',
          id: '1370458191',
          deposit: '已缴纳',
          rank: '白银'
        },
        {
          nickname: 'ZZ_yang',
          id: '2392005061',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '内裤假面蹲厕所（拉斯特炫）',
          id: '4217330893',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '情绪动物',
          id: '2280976518',
          deposit: '已缴纳',
          rank: '钻石'
        },
        {
          nickname: '伊卡洛斯',
          id: '2188958880',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '不如回家吃鸡翅',
          id: '3538824011',
          deposit: '已缴纳',
          rank: '黄金'
        },
        {
          nickname: 'JoJo',
          id: '1934089790',
          deposit: '已缴纳',
          rank: '钻石'
        },
        {
          nickname: '想成为超级升龙拳高手',
          id: '1305801683',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: '杰米爱喝酒',
          id: '3224382623',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: 'AAA不吃香菜的张斌',
          id: '3327401795',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: '西兰麻花',
          id: '2406625831',
          deposit: '已缴纳',
          rank: '黄金'
        },
        {
          nickname: 'Sakura02',
          id: '2354283833',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: '我要成为打差合高手',
          id: '1476349941',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '中二的大黑',
          id: '1254795536',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: '你比我更龙',
          id: '2545472433',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: '遇木木.',
          id: '4221833114',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: '渊渟岳峙',
          id: '2964463762',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'Ziyang',
          id: '1162911699',
          deposit: '已缴纳',
          rank: '钻石'
        },
        {
          nickname: '布洛芬缓释胶囊',
          id: '1097734803',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'zHnnnnnG',
          id: '4067095830',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: 'THE KING OF MAI',
          id: '2527493937',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: 'AstraStranger',
          id: '3717665799',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: '学外语好啊，外语得学',
          id: '4232499204',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: 'Forging',
          id: '1829183625',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: '梅子青时节',
          id: '1106590645',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: '宇哥薄纱小杰',
          id: '2158541130',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '无敌小塔',
          id: '1617883033',
          deposit: '已缴纳',
          rank: '钻石'
        },
        {
          nickname: 'N1ko',
          id: '3512735486',
          deposit: '已缴纳',
          rank: '白金'
        },
        {
          nickname: '三次回绝刘亦菲',
          id: '3446561559',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'ANDY',
          id: '1406520717',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '路易十二',
          id: '1626273883',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '|背手摔',
          id: '1561569085',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'windgrin',
          id: '2355954159',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'Rzez',
          id: '1541237644',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'Mark',
          id: '1895673609',
          deposit: '已缴纳',
          rank: '正分大师'
        },
        {
          nickname: 'M1rAgE',
          id: '4188692058',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'mocurrey',
          id: '1377614202',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'GDTT',
          id: '4080817483',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '永雏白给大王',
          id: '3070929526',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '打不过你你还羞辱我是什么意思',
          id: '3325554957',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: 'RYWWWWY',
          id: '1109544551',
          deposit: '已缴纳',
          rank: '黑铁'
        },
        {
          nickname: '权御天下',
          id: '1475709404',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '我便繁花似锦',
          id: '1284533117',
          deposit: '已缴纳',
          rank: '白银'
        },
        {
          nickname: '佐仓杏子',
          id: '2955418771',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '特瑞萌新FANG',
          id: '2953793950',
          deposit: '已缴纳',
          rank: '钻石'
        },
      ],
      round1Matches: [
        { id: 1, playerA: '历战街霸金狮子', playerB: 'RYWWWWY' },
        { id: 2, playerA: '我便繁花似锦', playerB: '长征永弓使' },
        { id: 3, playerA: '西兰麻花', playerB: '不如回家吃鸡翅' },
        { id: 4, playerA: 'N1ko', playerB: '梅子青时节' },
        { id: 5, playerA: 'Forging', playerB: 'THE KING OF MAI' },
        { id: 6, playerA: 'zHnnnnnG', playerB: '你比我更龙' },
        { id: 7, playerA: '中二的大黑', playerB: 'AAA不吃香菜的张斌' },
        { id: 8, playerA: '杰米爱喝酒', playerB: '胆小的阿烙' },
        { id: 9, playerA: 'Farid', playerB: '拉面馒头' },
        { id: 10, playerA: '特瑞萌新FANG', playerB: '无敌小塔' },
        { id: 11, playerA: 'Ziyang', playerB: 'JoJo' },
        { id: 12, playerA: 'Mark', playerB: 'AstraStranger' },
        { id: 13, playerA: '遇木木.', playerB: 'Sakura02' },
        { id: 14, playerA: '想成为超级升龙拳高手', playerB: '意墨要起飞' },
        { id: 15, playerA: '情绪动物', playerB: '佐仓杏子' },
        { id: 16, playerA: '一等情事', playerB: '权御天下' },
        { id: 17, playerA: '打不过你你还羞辱我是什么意思', playerB: '永雏白给大王' },
        { id: 18, playerA: 'GDTT', playerB: 'mocurrey' },
        { id: 19, playerA: 'M1rAgE', playerB: 'Rzez' },
        { id: 20, playerA: 'windgrin', playerB: '|背手摔' },
        { id: 21, playerA: '路易十二', playerB: 'ANDY' },
        { id: 22, playerA: '三次回绝刘亦菲', playerB: '宇哥薄纱小杰' },
        { id: 23, playerA: '布洛芬缓释胶囊', playerB: '渊渟岳峙' },
        { id: 24, playerA: '我要成为打差合高手', playerB: '伊卡洛斯' },
        { id: 25, playerA: '内裤假面蹲厕所（拉斯特炫）', playerB: 'ZZ_yang' },
        { id: 26, playerA: '我是一个超级大狗腿子', playerB: 'AzS_绕远路 才是最短的捷径' },
        { id: 27, playerA: '也许不是不开心', playerB: '鱼见千寻' },
        { id: 28, playerA: 'chilleroo', playerB: '我菜我不难受吗' },
        { id: 29, playerA: '左眼下起雨', playerB: 'Riven' },
        { id: 30, playerA: '两仪未来', playerB: '红色流浪星' },
        { id: 31, playerA: '若有聲聞者盡皆斬落斷魂於此', playerB: 'M别打了别打了M' },
        { id: 32, playerA: '我家猫会流氓后空翻夹头', playerB: 'linrichu' }
      ]
    };
  },
  computed: {
    reversedParticipants() {
      return [...this.participants].reverse();
    }
  }
};
</script>

<style scoped>
.registration-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
  min-height: 100vh;
  background-color: #fafafa;
}

.header {
  text-align: center;
  margin-bottom: 40px;
}

.header h1 {
  font-size: 32px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.announcement-wrapper {
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 30px;
  overflow: hidden;
}

.announcement-title {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #ffffff;
  font-size: 20px;
  font-weight: 600;
  padding: 16px 24px;
  text-align: center;
}

.announcement-content {
  padding: 24px;
}

.announcement-item {
  margin-bottom: 16px;
  line-height: 1.6;
  font-size: 15px;
  color: #333333;
}

.announcement-item:last-of-type {
  margin-bottom: 0;
}

.item-label {
  font-weight: 600;
  color: #2c3e50;
  display: inline-block;
  min-width: 100px;
}

.item-value {
  color: #555555;
}

.announcement-item strong {
  color: #667eea;
  font-weight: 600;
}

.announcement-note {
  margin-top: 20px;
  padding: 12px 16px;
  background-color: #f0f4ff;
  border-left: 4px solid #667eea;
  border-radius: 4px;
  font-size: 14px;
  color: #555555;
  line-height: 1.6;
}

.round1-wrapper {
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 30px;
  overflow: hidden;
}

.round1-title {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  color: #ffffff;
  font-size: 20px;
  font-weight: 600;
  padding: 16px 24px;
  text-align: center;
}

.round1-content {
  padding: 24px;
}

.round1-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 16px;
}

.round1-table thead {
  background-color: #f5f5f5;
}

.round1-table th {
  padding: 16px 20px;
  text-align: left;
  font-weight: 600;
  color: #2c3e50;
  border-bottom: 2px solid #e0e0e0;
}

.round1-table tbody tr {
  border-bottom: 1px solid #f0f0f0;
  transition: background-color 0.2s;
}

.round1-table tbody tr:hover {
  background-color: #f9f9f9;
}

.round1-table tbody tr:last-child {
  border-bottom: none;
}

.round1-table td {
  padding: 16px 20px;
  color: #333333;
}

.round1-table td:first-child {
  font-weight: 500;
}

.rules-wrapper {
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 30px;
  overflow: hidden;
}

.rules-title {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  color: #ffffff;
  font-size: 20px;
  font-weight: 600;
  padding: 16px 24px;
  text-align: center;
}

.rules-content {
  padding: 24px;
}

.rules-section {
  margin-bottom: 28px;
}

.rules-section:last-child {
  margin-bottom: 0;
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0 0 16px 0;
  padding-bottom: 8px;
  border-bottom: 2px solid #f0f0f0;
}

.section-content {
  color: #333333;
  line-height: 1.8;
  font-size: 15px;
}

.section-content p {
  margin: 8px 0;
}

.section-content p:first-child {
  margin-top: 0;
}

.section-content p:last-child {
  margin-bottom: 0;
}

.highlight-text {
  color: #f5576c;
  font-weight: 500;
  padding-left: 12px;
  border-left: 3px solid #f5576c;
  margin: 12px 0 !important;
}

.subsection {
  margin-bottom: 20px;
}

.subsection:last-child {
  margin-bottom: 0;
}

.subsection-title {
  font-size: 16px;
  font-weight: 600;
  color: #555555;
  margin: 0 0 12px 0;
}

.rules-list {
  margin: 12px 0;
  padding-left: 24px;
  list-style-type: none;
}

.rules-list li {
  margin-bottom: 10px;
  padding-left: 20px;
  position: relative;
  line-height: 1.7;
}

.rules-list li::before {
  content: "•";
  position: absolute;
  left: 0;
  color: #f5576c;
  font-weight: bold;
  font-size: 18px;
}

.rules-list li:last-child {
  margin-bottom: 0;
}

.rules-list li strong {
  color: #2c3e50;
  font-weight: 600;
}

.rules-wish {
  margin-top: 20px !important;
  text-align: center;
  font-size: 16px;
  font-weight: 500;
  color: #f5576c;
  padding: 12px;
  background-color: #fff5f7;
  border-radius: 6px;
}

.stats-wrapper {
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 30px;
  padding: 20px 24px;
}

.stats-content {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
}

.stats-label {
  color: #555555;
  font-weight: 500;
  margin-right: 8px;
}

.stats-value {
  color: #667eea;
  font-size: 28px;
  font-weight: 700;
  margin: 0 4px;
}

.stats-unit {
  color: #555555;
  font-weight: 500;
  margin-left: 4px;
}

.table-wrapper {
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.registration-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 16px;
}

.registration-table thead {
  background-color: #f5f5f5;
}

.registration-table th {
  padding: 16px 20px;
  text-align: left;
  font-weight: 600;
  color: #2c3e50;
  border-bottom: 2px solid #e0e0e0;
}

.registration-table tbody tr {
  border-bottom: 1px solid #f0f0f0;
  transition: background-color 0.2s;
}

.registration-table tbody tr:hover {
  background-color: #f9f9f9;
}

.registration-table tbody tr:last-child {
  border-bottom: none;
}

.registration-table td {
  padding: 16px 20px;
  color: #333333;
}

.registration-table td:first-child {
  font-weight: 500;
}

@media (max-width: 768px) {
  .registration-container {
    padding: 20px 10px;
  }

  .header h1 {
    font-size: 24px;
  }

  .announcement-title {
    font-size: 18px;
    padding: 12px 16px;
  }

  .announcement-content {
    padding: 16px;
  }

  .announcement-item {
    font-size: 14px;
    margin-bottom: 12px;
  }

  .item-label {
    display: block;
    margin-bottom: 4px;
    min-width: auto;
  }

  .announcement-note {
    font-size: 13px;
    padding: 10px 12px;
  }

  .round1-title {
    font-size: 18px;
    padding: 12px 16px;
  }

  .round1-content {
    padding: 16px;
  }

  .round1-table {
    font-size: 14px;
  }

  .round1-table th,
  .round1-table td {
    padding: 12px 10px;
  }

  .rules-title {
    font-size: 18px;
    padding: 12px 16px;
  }

  .rules-content {
    padding: 16px;
  }

  .rules-section {
    margin-bottom: 24px;
  }

  .section-title {
    font-size: 16px;
    margin-bottom: 12px;
  }

  .section-content {
    font-size: 14px;
  }

  .subsection-title {
    font-size: 15px;
    margin-bottom: 10px;
  }

  .rules-list {
    padding-left: 20px;
  }

  .rules-list li {
    padding-left: 16px;
    margin-bottom: 8px;
    font-size: 14px;
  }

  .rules-wish {
    font-size: 15px;
    padding: 10px;
  }

  .stats-wrapper {
    padding: 16px 20px;
    margin-bottom: 20px;
  }

  .stats-content {
    font-size: 16px;
  }

  .stats-value {
    font-size: 24px;
  }

  .registration-table {
    font-size: 14px;
  }

  .registration-table th,
  .registration-table td {
    padding: 12px 10px;
  }
}
</style>
