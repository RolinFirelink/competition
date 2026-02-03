<template>
  <div class="landing-page">
    <header class="hero">
      <div class="brand">Street Fighter 6 教学中心</div>
      <div class="hero-actions">
        <router-link class="pill active" to="/competition">赛事中心</router-link>
        <button class="pill disabled" disabled>攻略中心 · 敬请期待</button>
        <button class="pill disabled" disabled>直播与录播 · 敬请期待</button>
      </div>
      <p class="tagline">从入门到进阶，赛事、教学一站式体验</p>
    </header>

    <section class="ads-strip" aria-label="推广位">
      <div class="ads-track" :style="trackStyle">
        <div
          v-for="(ad, index) in adsLoop"
          :key="index"
          class="ad-item"
        >
          <a :href="ad.link" target="_blank" rel="noopener noreferrer">
            <img :src="ad.image" :alt="ad.title" />
            <span>{{ ad.title }}</span>
          </a>
        </div>
      </div>
    </section>

    <main class="content">
      <div class="section-title">快速入口</div>
      <div class="options-grid">
        <router-link class="option-card" to="/competition">
          <div class="option-label">赛事中心</div>
          <div class="option-desc">杯赛记录 · 报名 · 赛程</div>
        </router-link>
        <div class="option-card disabled">
          <div class="option-label">教学攻略</div>
          <div class="option-desc">角色教学 · 连招指令 · 对局心得</div>
        </div>
        <div class="option-card disabled">
          <div class="option-label">直播/录像</div>
          <div class="option-desc">赛事直播 · 高光集锦</div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  name: 'CompetitionLanding',
  data() {
    return {
      ads: [
        {
          title: '盖世小鸡 启明星手柄',
          link: 'https://www.gamestar.com',
          image:
            'https://dummyimage.com/320x120/1f1f1f/ffffff&text=Sponsor+1'
        },
        {
          title: '每日一灯 赛事直播',
          link: 'https://www.bilibili.com',
          image:
            'https://dummyimage.com/320x120/222222/ffffff&text=Sponsor+2'
        },
        {
          title: '街霸6 玩家社区',
          link: 'https://www.capcom.com',
          image:
            'https://dummyimage.com/320x120/262626/ffffff&text=Sponsor+3'
        }
      ],
      speed: 40
    };
  },
  computed: {
    adsLoop() {
      return [...this.ads, ...this.ads];
    },
    trackStyle() {
      const duration = (this.ads.length * 10) / (this.speed / 40);
      return {
        '--duration': `${duration}s`
      };
    }
  }
};
</script>

<style scoped>
.landing-page {
  min-height: 100vh;
  background: #0b0c10;
  color: #e8ecf1;
  font-family: 'Inter', 'Segoe UI', system-ui, -apple-system, sans-serif;
  padding: 48px 32px 64px;
}

.hero {
  max-width: 1200px;
  margin: 0 auto 32px auto;
  text-align: center;
}

.brand {
  font-size: 32px;
  font-weight: 800;
  letter-spacing: 0.5px;
  margin-bottom: 16px;
  color: #f5f7fb;
}

.hero-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 12px;
}

.pill {
  border: 1px solid #2a2d34;
  background: linear-gradient(135deg, #15171c 0%, #101117 100%);
  color: #9aa4b5;
  padding: 10px 18px;
  border-radius: 999px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.pill.active {
  border-color: #6c7bff;
  color: #dfe6ff;
  box-shadow: 0 0 0 1px rgba(108, 123, 255, 0.3);
}

.pill.disabled {
  cursor: not-allowed;
  opacity: 0.5;
}

.tagline {
  margin-top: 6px;
  color: #8f98ad;
  font-size: 15px;
}

.ads-strip {
  margin: 32px auto 40px;
  max-width: 1100px;
  overflow: hidden;
  border-radius: 14px;
  border: 1px solid #1c1f27;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.03), rgba(255, 255, 255, 0.01));
  padding: 12px;
}

.ads-track {
  display: flex;
  gap: 12px;
  animation: scroll var(--duration) linear infinite;
  width: max-content;
}

.ad-item {
  min-width: 320px;
  background: #12141a;
  border-radius: 10px;
  border: 1px solid #1f222c;
  overflow: hidden;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.25);
}

.ad-item a {
  display: flex;
  flex-direction: column;
  color: inherit;
  text-decoration: none;
}

.ad-item img {
  width: 100%;
  height: 120px;
  object-fit: cover;
  background: #0f1116;
}

.ad-item span {
  padding: 10px 12px;
  font-weight: 600;
  font-size: 14px;
}

@keyframes scroll {
  from {
    transform: translateX(0);
  }
  to {
    transform: translateX(calc(-50% - 6px));
  }
}

.content {
  max-width: 1100px;
  margin: 0 auto;
}

.section-title {
  font-size: 18px;
  font-weight: 700;
  color: #cfd5e3;
  margin-bottom: 14px;
}

.options-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 16px;
}

.option-card {
  background: radial-gradient(circle at 20% 20%, rgba(108, 123, 255, 0.15), transparent 35%),
    radial-gradient(circle at 80% 0%, rgba(255, 99, 164, 0.15), transparent 30%),
    #0f1116;
  border: 1px solid #1f222c;
  border-radius: 14px;
  padding: 18px;
  color: #e9edf5;
  text-decoration: none;
  display: flex;
  flex-direction: column;
  gap: 8px;
  transition: transform 0.15s ease, border-color 0.2s ease;
}

.option-card:hover {
  transform: translateY(-4px);
  border-color: #6c7bff;
}

.option-card.disabled {
  cursor: not-allowed;
  opacity: 0.5;
}

.option-label {
  font-size: 18px;
  font-weight: 800;
  letter-spacing: 0.3px;
}

.option-desc {
  font-size: 13px;
  color: #9aa4b5;
}

@media (max-width: 720px) {
  .landing-page {
    padding: 32px 20px 48px;
  }

  .brand {
    font-size: 26px;
  }

  .ad-item {
    min-width: 260px;
  }
}
</style>

