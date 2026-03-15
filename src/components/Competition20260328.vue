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
            <td>{{ participant.rank }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Competition20260328',
  data() {
    return {
      qrCodeUrl:
        'https://rolin-typora.oss-cn-guangzhou.aliyuncs.com/91faac1245645b3b0c31a69f80b06993.jpg',
      showQrPreview: false,
      participants: [
        {
          nickname: '示例选手',
          id: '0000000000',
          rank: '待定'
        }
      ]
    };
  },
  methods: {
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
</style>


