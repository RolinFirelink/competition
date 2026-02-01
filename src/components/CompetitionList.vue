<template>
  <div class="competition-list-container">
    <div class="hero-section">
      <div class="header">
        <h1>🏆 赛事中心</h1>
        <p class="subtitle">Street Fighter 6 杯赛记录</p>
      </div>
    </div>
    <div class="competitions-wrapper">
      <div class="competition-item" v-for="competition in competitions" :key="competition.path">
        <router-link :to="competition.path" class="competition-link">
          <div class="competition-header">
            <div class="competition-date">{{ competition.date }}</div>
            <div class="competition-badge">已结束</div>
          </div>
          <div class="competition-title">{{ competition.title }}</div>
          <div class="competition-info">
            <div class="info-item">
              <span class="info-icon">👥</span>
              <span class="info-text">参赛人数：<strong>{{ competition.participants }}</strong>人</span>
            </div>
          </div>
          <div class="sponsor-section">
            <div class="sponsor-item sponsor-company">
              <div class="sponsor-label">
                <span class="sponsor-icon">🏢</span>
                <span>赞助商</span>
              </div>
              <div class="sponsor-name">{{ competition.sponsor }}</div>
            </div>
            <div class="sponsor-item sponsor-person">
              <div class="sponsor-label">
                <span class="sponsor-icon">🙏</span>
                <span>赞助人</span>
              </div>
              <div class="sponsor-name">{{ competition.patron }}</div>
            </div>
          </div>
          <div class="view-details">
            <span>查看详情 →</span>
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CompetitionList',
  data() {
    return {
      competitions: [
        {
          path: '/20260131',
          date: '2026年1月31日',
          title: '2026年1月31日杯赛',
          participants: 70,
          sponsor: '盖世小鸡',
          patron: 'H半瓶'
        }
        // 未来可以在这里添加更多赛事
      ]
    };
  }
};
</script>

<style scoped>
.competition-list-container {
  max-width: 1400px;
  margin: 0 auto;
  padding: 0;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  position: relative;
}

.competition-list-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: 
    radial-gradient(circle at 20% 50%, rgba(255, 255, 255, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 80% 80%, rgba(255, 255, 255, 0.1) 0%, transparent 50%);
  pointer-events: none;
}

.hero-section {
  padding: 60px 20px 40px;
  text-align: center;
  position: relative;
  z-index: 1;
}

.header h1 {
  font-size: 48px;
  font-weight: 800;
  color: #ffffff;
  margin: 0 0 12px 0;
  text-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  letter-spacing: -1px;
}

.subtitle {
  font-size: 18px;
  color: rgba(255, 255, 255, 0.9);
  margin: 0;
  font-weight: 400;
}

.competitions-wrapper {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(380px, 1fr));
  gap: 32px;
  padding: 0 40px 60px;
  position: relative;
  z-index: 1;
}

.competition-item {
  background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.2);
  position: relative;
}

.competition-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
}

.competition-item:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 16px 48px rgba(0, 0, 0, 0.25);
}

.competition-link {
  display: block;
  padding: 32px;
  text-decoration: none;
  color: inherit;
  position: relative;
}

.competition-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.competition-date {
  font-size: 13px;
  color: #667eea;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 1.5px;
  padding: 6px 12px;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 20px;
  display: inline-block;
}

.competition-badge {
  font-size: 11px;
  color: #ffffff;
  font-weight: 600;
  padding: 4px 12px;
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  border-radius: 12px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.competition-title {
  font-size: 28px;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 24px;
  line-height: 1.3;
}

.competition-info {
  margin-bottom: 24px;
  padding-bottom: 24px;
  border-bottom: 2px solid #f0f0f0;
}

.info-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 15px;
  color: #555555;
}

.info-icon {
  font-size: 18px;
}

.info-text strong {
  color: #667eea;
  font-weight: 700;
  font-size: 16px;
}

.sponsor-section {
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin-bottom: 24px;
  padding: 20px;
  background: linear-gradient(135deg, #fff9e6 0%, #fff5f7 100%);
  border-radius: 16px;
  border: 2px solid #ffd700;
  box-shadow: 0 4px 16px rgba(255, 215, 0, 0.15);
}

.sponsor-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #ffffff;
  border-radius: 12px;
  transition: transform 0.2s ease;
}

.sponsor-item:hover {
  transform: translateX(4px);
}

.sponsor-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #666666;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.sponsor-icon {
  font-size: 18px;
}

.sponsor-name {
  font-size: 20px;
  font-weight: 800;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  letter-spacing: 0.5px;
}

.sponsor-company .sponsor-name {
  background: linear-gradient(135deg, #4a90e2 0%, #357abd 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.sponsor-person .sponsor-name {
  background: linear-gradient(135deg, #f39c12 0%, #e67e22 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.view-details {
  text-align: center;
  padding-top: 20px;
  border-top: 1px solid #f0f0f0;
  font-size: 14px;
  font-weight: 600;
  color: #667eea;
  transition: color 0.3s ease;
}

.competition-item:hover .view-details {
  color: #764ba2;
}

.view-details span {
  display: inline-block;
  transition: transform 0.3s ease;
}

.competition-item:hover .view-details span {
  transform: translateX(4px);
}

@media (max-width: 768px) {
  .competition-list-container {
    padding: 0;
  }

  .hero-section {
    padding: 40px 20px 30px;
  }

  .header h1 {
    font-size: 32px;
  }

  .subtitle {
    font-size: 16px;
  }

  .competitions-wrapper {
    grid-template-columns: 1fr;
    gap: 24px;
    padding: 0 20px 40px;
  }

  .competition-link {
    padding: 24px;
  }

  .competition-title {
    font-size: 22px;
  }

  .sponsor-section {
    padding: 16px;
  }

  .sponsor-name {
    font-size: 18px;
  }
}
</style>

