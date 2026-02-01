<template>
  <div class="registration-container">
    <div class="header">
      <h1>2026年1月31日杯赛报名情况</h1>
    </div>
    <div class="announcement-wrapper">
      <div class="announcement-title">📢 赛事公告</div>
      <div class="announcement-content">
        <div class="announcement-item">
          <span class="item-label">开赛时间：</span>
          <span class="item-value">2026年1月31日下午两点</span>
        </div>
        <div class="announcement-item">
          <span class="item-label">报名方式：</span>
          <span class="item-value">请加QQ群 <strong>743749814</strong></span>
        </div>
        <div class="announcement-item">
          <span class="item-label">保证金说明：</span>
          <span class="item-value">报名需缴纳10元保证金。正常参赛将原路退回；如鸽子，保证金将作为奖励给到随机一位选手。</span>
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
    <div class="replay-wrapper">
      <div class="replay-title">📺 赛事全程回放</div>
      <div class="replay-content">
        <a :href="replayUrl" target="_blank" class="replay-link" rel="noopener noreferrer">
          <span class="replay-icon">▶️</span>
          <span class="replay-text">点击观看B站全程回放</span>
          <span class="replay-arrow">→</span>
        </a>
      </div>
    </div>
    <div class="champion-wrapper">
      <div class="champion-title">🏆 冠军</div>
      <div class="champion-content">
        <div class="champion-name">{{ champion }}</div>
      </div>
    </div>
    <div class="awards-wrapper">
      <div class="award-item">
        <div class="award-icon">🎮</div>
        <div class="award-info">
          <div class="award-title">盖世小鸡启明星手柄幸运奖</div>
          <div class="award-winner">{{ luckyPrizeWinner }}</div>
        </div>
      </div>
      <div class="award-item">
        <div class="award-icon">🎁</div>
        <div class="award-info">
          <div class="award-title">幸运金奖</div>
          <div class="award-winner">{{ luckyGoldWinner }}</div>
        </div>
      </div>
    </div>
    <div class="round1-wrapper">
      <div class="round1-title">📋 决赛名单</div>
      <div class="round1-content">
        <div class="round1-note">
          💡 每组分批打，从上往下，第一组打完轮到下一组，每个人有三分钟的迟到时间，迟到算弃赛，特殊情况延后到五分钟。
        </div>
        <table class="round1-table">
          <thead>
            <tr>
              <th>序号</th>
              <th>昵称A</th>
              <th>昵称B</th>
            </tr>
          </thead>
          <tbody>
            <!-- 第一组 -->
            <tr class="group-row">
              <td colspan="3">第一组</td>
            </tr>
            <tr v-for="match in finalMatches.slice(0,4)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="round1-wrapper">
      <div class="round1-title">📋 四强赛名单</div>
      <div class="round1-content">
        <div class="round1-note">
          💡 每组分批打，从上往下，第一组打完轮到下一组，每个人有三分钟的迟到时间，迟到算弃赛，特殊情况延后到五分钟。
        </div>
        <table class="round1-table">
          <thead>
            <tr>
              <th>序号</th>
              <th>昵称A</th>
              <th>昵称B</th>
            </tr>
          </thead>
          <tbody>
            <!-- 第一组 -->
            <tr class="group-row">
              <td colspan="3">第一组</td>
            </tr>
            <tr v-for="match in semiFinalsMatches.slice(0,4)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="round1-wrapper">
      <div class="round1-title">📋 八强赛名单</div>
      <div class="round1-content">
        <div class="round1-note">
          💡 每组分批打，从上往下，第一组打完轮到下一组，每个人有三分钟的迟到时间，迟到算弃赛，特殊情况延后到五分钟。
        </div>
        <table class="round1-table">
          <thead>
            <tr>
              <th>序号</th>
              <th>昵称A</th>
              <th>昵称B</th>
            </tr>
          </thead>
          <tbody>
            <!-- 第一组 -->
            <tr class="group-row">
              <td colspan="3">第一组</td>
            </tr>
            <tr v-for="match in quarterFinalsMatches.slice(0,4)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="round1-wrapper">
      <div class="round1-title">📋 淘汰赛第四轮名单</div>
      <div class="round1-content">
        <div class="round1-note">
          💡 每组分批打，从上往下，第一组打完轮到下一组，每个人有三分钟的迟到时间，迟到算弃赛，特殊情况延后到五分钟。
        </div>
        <table class="round1-table">
          <thead>
            <tr>
              <th>序号</th>
              <th>昵称A</th>
              <th>昵称B</th>
            </tr>
          </thead>
          <tbody>
            <!-- 第一组 -->
            <tr class="group-row">
              <td colspan="3">第一组</td>
            </tr>
            <tr v-for="match in round4Matches.slice(0,4)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="round1-wrapper">
      <div class="round1-title">📋 淘汰赛第三轮名单</div>
      <div class="round1-content">
        <div class="round1-note">
          💡 每组分批打，从上往下，第一组打完轮到下一组，每个人有三分钟的迟到时间，迟到算弃赛，特殊情况延后到五分钟。
        </div>
        <table class="round1-table">
          <thead>
            <tr>
              <th>序号</th>
              <th>昵称A</th>
              <th>昵称B</th>
            </tr>
          </thead>
          <tbody>
            <!-- 第一组 -->
            <tr class="group-row">
              <td colspan="3">第一组</td>
            </tr>
            <tr v-for="match in round3Matches.slice(0,4)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第二组 -->
            <tr class="group-row">
              <td colspan="3">第二组</td>
            </tr>
            <tr v-for="match in round3Matches.slice(4,8)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="round1-wrapper">
      <div class="round1-title">📋 淘汰赛第二轮名单</div>
      <div class="round1-content">
        <div class="round1-note">
          💡 每组分批打，从上往下，第一组打完轮到下一组，每个人有三分钟的迟到时间，迟到算弃赛，特殊情况延后到五分钟。
        </div>
        <table class="round1-table">
          <thead>
            <tr>
              <th>序号</th>
              <th>昵称A</th>
              <th>昵称B</th>
            </tr>
          </thead>
          <tbody>
            <!-- 第一组 -->
            <tr class="group-row">
              <td colspan="3">第一组</td>
            </tr>
            <tr v-for="match in round2Matches.slice(0,4)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第二组 -->
            <tr class="group-row">
              <td colspan="3">第二组</td>
            </tr>
            <tr v-for="match in round2Matches.slice(4,8)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第三组 -->
            <tr class="group-row">
              <td colspan="3">第三组</td>
            </tr>
            <tr v-for="match in round2Matches.slice(8,12)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第四组 -->
            <tr class="group-row">
              <td colspan="3">第四组</td>
            </tr>
            <tr v-for="match in round2Matches.slice(12,16)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="round1-wrapper">
      <div class="round1-title">📋 淘汰赛第一轮名单</div>
      <div class="round1-content">
        <div class="round1-note">
          💡 每组分批打，从上往下，第一组打完轮到下一组，每个人有三分钟的迟到时间，迟到算弃赛，特殊情况延后到五分钟。
        </div>
        <table class="round1-table">
          <thead>
            <tr>
              <th>序号</th>
              <th>昵称A</th>
              <th>昵称B</th>
            </tr>
          </thead>
          <tbody>
            <!-- 第0组 -->
            <tr class="group-row">
              <td colspan="3">第0组</td>
            </tr>
            <tr>
              <td>0</td>
              <td>刘康x嘉米是不是邪门CP</td>
              <td>领导，我干了，你随意</td>
            </tr>
            <!-- 第一组 -->
            <tr class="group-row">
              <td colspan="3">第一组</td>
            </tr>
            <tr v-for="match in round1Matches.slice(0,4)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第二组 -->
            <tr class="group-row">
              <td colspan="3">第二组</td>
            </tr>
            <tr v-for="match in round1Matches.slice(4,8)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第三组 -->
            <tr class="group-row">
              <td colspan="3">第三组</td>
            </tr>
            <tr v-for="match in round1Matches.slice(8,12)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第四组 -->
            <tr class="group-row">
              <td colspan="3">第四组</td>
            </tr>
            <tr v-for="match in round1Matches.slice(12,16)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第五组 -->
            <tr class="group-row">
              <td colspan="3">第五组</td>
            </tr>
            <tr v-for="match in round1Matches.slice(16,20)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第六组 -->
            <tr class="group-row">
              <td colspan="3">第六组</td>
            </tr>
            <tr v-for="match in round1Matches.slice(20,24)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第七组 -->
            <tr class="group-row">
              <td colspan="3">第七组</td>
            </tr>
            <tr v-for="match in round1Matches.slice(24,28)" :key="match.id">
              <td>{{ match.id }}</td>
              <td>{{ match.playerA }}</td>
              <td>{{ match.playerB }}</td>
            </tr>
            <!-- 第八组 -->
            <tr class="group-row">
              <td colspan="3">第八组</td>
            </tr>
            <tr v-for="match in round1Matches.slice(28,32)" :key="match.id">
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
              <li><strong>签到与迟到：</strong>每场开始前会在群内通知，选手有3分钟准备时间。超时无理由视为弃赛，对手晋级，保证金计入奖池。特殊情况可延至5分钟。</li>
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
          deposit: '已入池',
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
          deposit: '已入池',
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
        {
          nickname: '领导，我干了，你随意',
          id: '1683485597',
          deposit: '已缴纳',
          rank: '负分大师'
        },
        {
          nickname: '刘康x嘉米算不算邪门CP?',
          id: '4239997827',
          deposit: '已免费',
          rank: '负分大师'
        },
      ],
      champion: '学外语好啊外语得学',
      replayUrl: 'https://www.bilibili.com/video/BV15R6JBUEyv/?spm_id_from=333.337.search-card.all.click&vd_source=eb48b40a0ac8610fb2a9c4b62b48f570',
      luckyPrizeWinner: '永雏白给大王',
      luckyGoldWinner: 'mocurrey',
      finalMatches: [
        { id: 1, playerA: '学外语好啊外语得学', playerB: 'SiegeCascade' }
      ],
      semiFinalsMatches: [
        { id: 1, playerA: 'Sakura02', playerB: '学外语好啊外语得学' },
        { id: 2, playerA: 'AOZaki', playerB: 'SiegeCascade' }
      ],
      quarterFinalsMatches: [
        { id: 1, playerA: '特瑞萌新FANG', playerB: 'SiegeCascade' },
        { id: 2, playerA: '左眼下起雨', playerB: 'AOZaki' },
        { id: 3, playerA: 'Sakura02', playerB: '喜欢罐罐' },
        { id: 4, playerA: '领导我干了你随意', playerB: '学外语好啊外语得学' }
      ],
      round4Matches: [
        { id: 1, playerA: '左眼下起雨', playerB: 'GDTT' },
        { id: 2, playerA: '你比我更龙', playerB: '特瑞萌新FANG' },
        { id: 3, playerA: '宇哥薄纱小杰', playerB: 'Sakura02' },
        { id: 4, playerA: '领导，我干了，你随意', playerB: '若有声闻者' }
      ],
      round3Matches: [
        { id: 1, playerA: '长征永弓使', playerB: '梅子青时节' },
        { id: 2, playerA: '你比我更龙', playerB: '杰米爱喝酒' },
        { id: 3, playerA: '特瑞萌新FANG', playerB: '佐仓信子' },
        { id: 4, playerA: 'GDTT', playerB: 'windgrin' },
        { id: 5, playerA: '宇哥薄纱小杰', playerB: '布洛芬缓释胶囊' },
        { id: 6, playerA: '一等情事', playerB: 'Sakura02' },
        { id: 7, playerA: '鱼见千寻', playerB: '左眼下起雨' },
        { id: 8, playerA: '我是一个超级大狗腿子', playerB: '若有声闻者' }
      ],
      round2Matches: [
        { id: 1, playerA: 'rywwwwy', playerB: '长征' },
        { id: 2, playerA: '西兰麻花', playerB: '梅子青时节' },
        { id: 3, playerA: 'the king of mai', playerB: '你比我更龙' },
        { id: 4, playerA: 'AAA不吃香菜', playerB: '杰米爱喝酒' },
        { id: 5, playerA: 'Farid', playerB: '特瑞萌新FANG' },
        { id: 6, playerA: 'JoJo', playerB: '佐仓信子' },
        { id: 7, playerA: '永雏白给大王', playerB: 'GDTT' },
        { id: 8, playerA: 'M1rage', playerB: 'windgrin' },
        { id: 9, playerA: '路易十二', playerB: '宇哥薄纱小杰' },
        { id: 10, playerA: '布洛芬缓释胶囊', playerB: '我要成为打差合高手' },
        { id: 11, playerA: '想成为超级升龙拳高手', playerB: '一等情事' },
        { id: 12, playerA: '鱼见千寻', playerB: '我菜我不难受吗' },
        { id: 13, playerA: '左眼下起雨', playerB: '红色流浪行' },
        { id: 14, playerA: '若有声闻者', playerB: 'linrichu' },
        { id: 15, playerA: 'AstraStranger', playerB: 'Sakura02' },
        { id: 16, playerA: '拉斯特炫', playerB: '我是一个超级大狗腿子' }
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
@import '../assets/styles/competition.css';
</style>