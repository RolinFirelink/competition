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
          </tr>
        </thead>
        <tbody>
          <tr v-for="(participant, index) in participants" :key="index">
            <td>{{ participant.nickname }}</td>
            <td>{{ participant.id }}</td>
            <td>{{ participant.deposit }}</td>
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
          deposit: '已缴纳'
        },
        {
          nickname: '我家猫会流氓后空翻夹头',
          id: '2783676023',
          deposit: '已缴纳'
        },
        {
          nickname: 'M别打了别打了M',
          id: '3272404266',
          deposit: '已缴纳'
        },
        {
          nickname: '若有聲聞者盡皆斬落斷魂於此',
          id: '3576414025',
          deposit: '已免费'
        },
        {
          nickname: '红色流浪星',
          id: '3498125113',
          deposit: '已缴纳'
        },
        {
          nickname: '胆小的阿烙',
          id: '3538353124',
          deposit: '已缴纳'
        },
        {
          nickname: '两仪未来',
          id: '3876379621',
          deposit: '已缴纳'
        },
        {
          nickname: 'AOZaki',
          id: '3431187004',
          deposit: '已缴纳'
        },
        {
          nickname: 'Riven',
          id: '2696437313',
          deposit: '已缴纳'
        },
        {
          nickname: '左眼下起雨',
          id: '4226859931',
          deposit: '已缴纳'
        },
        {
          nickname: '我菜我不难受吗',
          id: '3766826973',
          deposit: '已缴纳'
        },
        {
          nickname: '拉面馒头',
          id: '3753462047',
          deposit: '已缴纳'
        },
        {
          nickname: '历战街霸金狮子',
          id: '2662803340',
          deposit: '已缴纳'
        },
        {
          nickname: 'chilleroo',
          id: '1897981597',
          deposit: '已缴纳'
        },
        {
          nickname: '鱼见千寻',
          id: '2177871520',
          deposit: '已缴纳'
        },
        {
          nickname: '也许不是不开心',
          id: '3791128655',
          deposit: '已缴纳'
        },
        {
          nickname: 'Farid',
          id: '4282409574',
          deposit: '已缴纳'
        },
        {
          nickname: 'AzS_绕远路 才是最短的捷径',
          id: '1315717156',
          deposit: '已缴纳'
        },
        {
          nickname: '我是一个超级大狗腿子',
          id: '1314720339',
          deposit: '已缴纳'
        },
        {
          nickname: '胆小的阿烙',
          id: '3538353124',
          deposit: '已缴纳'
        },
        {
          nickname: '喜欢罐罐',
          id: '1324496630',
          deposit: '已缴纳'
        },
        {
          nickname: '意墨要起飞',
          id: '2171623952',
          deposit: '已缴纳'
        },
        {
          nickname: 'SiegeCascade',
          id: '1083334838',
          deposit: '已缴纳'
        },
        {
          nickname: '长征永弓使',
          id: '1370458191',
          deposit: '已缴纳'
        },
        {
          nickname: 'ZZ_yang',
          id: '2392005061',
          deposit: '已缴纳'
        },
        {
          nickname: ':内裤假面蹲厕所（拉斯特炫）',
          id: '4217330893',
          deposit: '已缴纳'
        },
      ]
    }
  }
}
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

