<template>
  <div class="registration-container">
    <div class="header">
      <h1>2026年3月28日 杯赛报名情况</h1>
    </div>

    <div class="announcement-wrapper">
      <div class="announcement-title">📢 赛事公告</div>
      <div class="announcement-content">
        <div class="announcement-item">
          <span class="item-label">开赛时间：</span>
          <span class="item-value">2026年3月28日 20:00</span>
        </div>
        <div class="announcement-item">
          <span class="item-label">报名方式：</span>
          <span class="item-value">
            扫描下方二维码进交流群完成报名
          </span>
        </div>
        <div class="announcement-item">
          <span class="item-label">赛事赞助：</span>
          <span class="item-value">
            赞助商：<strong>HJK</strong><br />
            特别鸣谢：<strong>HJK官方</strong>、天手力、Ohhhhhua、月光熔铁星、江舞、中二的大黑、阿依在
          </span>
        </div>
        <div class="announcement-item">
          <span class="item-label">赛制：</span>
          <span class="item-value">
            八强之前抢三，八强及之后都是抢五，不设复活赛。由于要凑2的幂,所以第一轮有17人随机轮空,剩余选手随机配对,第二轮将会由第一轮晋级选手与轮空选手配对对战。
            第三轮及之后将一路往上打直到决赛。
          </span>
        </div>
        <div class="announcement-item">
          <span class="item-label">声明：</span>
          <span class="item-value">
            八强之前并行打，八强之后一对一对的形式上场打。
          </span>
        </div>

        <div class="announcement-item qr-section">
          <span class="item-label">报名二维码：</span>
          <div class="qr-wrapper">
            <img
              class="qr-image"
              :src="qrCodeUrl"
              alt="报名二维码"
              @click="openQrPreview"
            />
            <div class="qr-tip">点击二维码查看大图</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 二维码放大预览 -->
    <div
      v-if="showQrPreview"
      class="qr-modal-overlay"
      @click="closeQrPreview"
    >
      <div class="qr-modal-content" @click.stop>
        <img :src="qrCodeUrl" alt="报名二维码大图" class="qr-modal-image" />
        <button class="qr-modal-close" @click="closeQrPreview">关闭</button>
      </div>
    </div>

    <div class="stats-wrapper">
      <div class="stats-content">
        <span class="stats-label">参赛人数：</span>
        <span class="stats-value">{{ participants.length }}</span>
        <span class="stats-unit">人</span>
      </div>
    </div>

    <div class="round1-wrapper">
      <div class="round1-title">📋 第一轮分组名单</div>
      <div class="round1-content">
        <div class="round1-note">
          💡 本轮随机抽取 <strong>17</strong> 人轮空；其余选手随机两两配对。每组最多 4 场（8人），不足部分自然减少场次。
        </div>

        <div class="round1-groups-grid">
          <div v-for="group in firstRoundGroups" :key="group.index" class="round1-group-card">
            <table class="round1-table">
              <thead>
                <tr>
                  <th>序号</th>
                  <th>昵称A</th>
                  <th>昵称B</th>
                </tr>
              </thead>
              <tbody>
                <tr class="group-row">
                  <td colspan="3">第{{ group.index }}组</td>
                </tr>
                <tr v-for="match in group.matches" :key="match.id">
                  <td>{{ match.id }}</td>
                  <td>{{ match.playerA.nickname }}</td>
                  <td>{{ match.playerB.nickname }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div class="bye-table-wrapper">
          <table class="registration-table bye-table">
            <thead>
              <tr>
                <th>序号</th>
                <th>轮空昵称</th>
                <th>街霸ID</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(p, idx) in byeParticipants" :key="idx">
                <td>{{ idx + 1 }}</td>
                <td>{{ p.nickname }}</td>
                <td>{{ p.id || '-' }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <div class="table-wrapper">
      <table class="registration-table">
        <thead>
          <tr>
            <th>街霸昵称</th>
            <th>街霸ID</th>
            <th>段位</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(participant, index) in participants" :key="index">
            <td>{{ participant.nickname }}</td>
            <td>{{ participant.id }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Competition20260328',
  created() {
    this.initFirstRound();
  },
  data() {
    return {
      qrCodeUrl:
        'https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/91faac1245645b3b0c31a69f80b06993.jpg',
      showQrPreview: false,
      byeParticipants: [],
      firstRoundGroups: [],
      participants: [
        {
          nickname: 'NPC哒',
          id: '2840493464',
          rank: '待定'
        },
        {
          nickname: 'Robin',
          id: '4078352348',
          rank: '待定'
        },
        {
          nickname: 'Xuers|艾笠如初',
          id: '1924890682',
          rank: '待定'
        },
        {
          nickname: '泼墨写春秋',
          id: '1713524530',
          rank: '待定'
        },
        {
          nickname: 'xuers｜土间大平',
          id: '3138934808',
          rank: '待定'
        },
        {
          nickname: '不想取名字の我',
          id: '2126980171',
          rank: '待定'
        },
        {
          nickname: 'hungry bird',
          id: '2151157324',
          rank: '待定'
        },
        {
          nickname: '米奇妙妙鬼',
          id: '1751741531',
          rank: '待定'
        },
        {
          nickname: '奈欧Neo',
          id: '1477253645',
          rank: '待定'
        },
        {
          nickname: 'WildFree.Stugx',
          id: '1616451996',
          rank: '待定'
        },
        {
          nickname: '南风南风',
          id: '1751414745',
          rank: '待定'
        },
        {
          nickname: 'Avid',
          id: '1437624290',
          rank: '待定'
        },
        {
          nickname: '格斗少女江玉燕',
          id: '2796482837',
          rank: '待定'
        },
        {
          nickname: 'Bilibili-昕缘刀妹',
          id: '2610472011',
          rank: '待定'
        },
        {
          nickname: 'FLin7',
          id: '1669229188',
          rank: '待定'
        },
        {
          nickname: 'ADO1337',
          id: '2113227374',
          rank: '待定'
        },
        {
          nickname: 'AOZaki',
          id: '3431187004',
          rank: '待定'
        },
        {
          nickname: 'Xuers|Yukiizh',
          id: '1192893558',
          rank: '待定'
        },
        {
          nickname: '月落有雪',
          id: '',
          rank: '待定'
        },
        {
          nickname: '2sk-超究极混沌暗影之狼狗',
          id: '3530046468',
          rank: '待定'
        },
        {
          nickname: '2sk-消失王',
          id: '4181392148',
          rank: '待定'
        },
        {
          nickname: 'GTW-小宅',
          id: '',
          rank: '待定'
        },
        {
          nickname: '学不到猛虎下山了',
          id: '3085727087',
          rank: '待定'
        },
        {
          nickname: 'HOT|丹阳大鸟',
          id: '3291211435',
          rank: '待定'
        },
        {
          nickname: '脏脏包',
          id: '4252431000',
          rank: '待定'
        },
        {
          nickname: '锤锤锤锤锤',
          id: '1175965295',
          rank: '待定'
        },
        {
          nickname: 'GTW-菜',
          id: '2949053681',
          rank: '待定'
        },
        {
          nickname: 'HOT I KLEM',
          id: '',
          rank: '待定'
        },
        {
          nickname: 'HJK l pinkstar',
          id: '2388112059',
          rank: '待定'
        },
        {
          nickname: '咲叶pandasaku',
          id: '',
          rank: '待定'
        },
        {
          nickname: '5hk领域大神',
          id: '3918309596',
          rank: '待定'
        },
        {
          nickname: 'COPE',
          id: '2450621083',
          rank: '待定'
        },
        {
          nickname: 'Dracula-K',
          id: '3559764253',
          rank: '待定'
        },
        {
          nickname: '33prs',
          id: '2165191875',
          rank: '待定'
        },
        {
          nickname: '-艾柯耀-',
          id: '2500339497',
          rank: '待定'
        },
        {
          nickname: '2sk-孤独的楼',
          id: '4090019287',
          rank: '待定'
        },
        {
          nickname: '我欲乘风归去',
          id: '1395398157',
          rank: '待定'
        },
        {
          nickname: '汉堡汉堡',
          id: '',
          rank: '待定'
        },
        {
          nickname: 'kuroasing',
          id: '1840204688',
          rank: '待定'
        },
        {
          nickname: 'GTW-XIAOHU',
          id: '1687147839',
          rank: '待定'
        },
        {
          nickname: '武汉洋洋',
          id: '3332414353',
          rank: '待定'
        },
        {
          nickname: 'GTW-XiaoXu',
          id: '3140819469',
          rank: '待定'
        },
        {
          nickname: '残荷听雨',
          id: '1839452670',
          rank: '待定'
        },
        {
          nickname: '名字真难取',
          id: '3091629655',
          rank: '待定'
        },
        {
          nickname: '我的精神有点问题',
          id: '1803160287',
          rank: '待定'
        },
        {
          nickname: '伤心公主',
          id: '3778650570',
          rank: '待定'
        },
        {
          nickname: '马嘉祺超绝肌肉线条',
          id: '3966050394',
          rank: '待定'
        },
      ]
    };
  },
  methods: {
    initFirstRound() {
      const byeCount = 17;
      const all = [...this.participants];
      const shuffled = this.shuffleArray(all);

      const byes = shuffled.slice(0, byeCount);
      const playPlayers = shuffled.slice(byeCount);

      const matches = [];
      for (let i = 0; i < playPlayers.length; i += 2) {
        const playerA = playPlayers[i];
        const playerB = playPlayers[i + 1];
        if (!playerB) break;
        matches.push({
          id: matches.length + 1,
          playerA,
          playerB
        });
      }

      const groups = [];
      const matchesPerGroup = 4; // 4场 = 8人
      for (let i = 0; i < matches.length; i += matchesPerGroup) {
        groups.push({
          index: groups.length + 1,
          matches: matches.slice(i, i + matchesPerGroup)
        });
      }

      this.byeParticipants = byes;
      this.firstRoundGroups = groups;
    },
    shuffleArray(arr) {
      // Fisher-Yates shuffle (in-place)
      for (let i = arr.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [arr[i], arr[j]] = [arr[j], arr[i]];
      }
      return arr;
    },
    openQrPreview() {
      this.showQrPreview = true;
    },
    closeQrPreview() {
      this.showQrPreview = false;
    }
  }
};
</script>

<style scoped>
@import '../assets/styles/competition.css';

.qr-section {
  align-items: flex-start;
}

.qr-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.qr-image {
  width: 160px;
  height: 160px;
  object-fit: cover;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
  cursor: pointer;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.qr-image:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.25);
}

.qr-tip {
  font-size: 12px;
  color: #888;
}

.qr-modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.qr-modal-content {
  background: #fff;
  padding: 16px 16px 20px;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
  text-align: center;
  max-width: 90vw;
  max-height: 90vh;
}

.qr-modal-image {
  max-width: 70vw;
  max-height: 70vh;
  border-radius: 8px;
}

.qr-modal-close {
  margin-top: 12px;
  padding: 6px 16px;
  border: none;
  border-radius: 999px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  font-size: 13px;
  cursor: pointer;
}

.round1-groups-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(360px, 1fr));
  gap: 16px;
  margin-top: 16px;
}

.round1-group-card {
  background: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  overflow: hidden;
}

.bye-table-wrapper {
  margin-top: 18px;
}

.bye-table {
  font-size: 15px;
}
</style>


