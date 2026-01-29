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
