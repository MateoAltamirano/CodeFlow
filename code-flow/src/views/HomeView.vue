<template>
  <div class="home">
    <div class="welcome">
      <h1 style="margin-left: 2rem">Welcome</h1>
      <div class="intro">
        <div class="intro-det">
          <h2>Code</h2>
          <img
            class="intro-img"
            style="object-fit: contain"
            src="../assets/Code.png"
            alt=""
          />
        </div>
        <div class="intro-det" style="align-items: flex-end">
          <img class="intro-img" src="../assets/Flow.png" alt="" />
          <h2>Flow</h2>
        </div>
      </div>
    </div>
    <div class="features">
      <div class="feature">
        <div class="feature-icon"></div>
        <div class="feature-det">
          <p style="font-size: 2rem">Code generator</p>
          <p>Transform your diagrams into programs and run them!</p>
        </div>
      </div>
      <div class="feature">
        <div class="feature-icon"></div>
        <div class="feature-det">
          <p style="font-size: 2rem">Node editor</p>
          <p>Drag and drop nodes to create simple programs!</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.home {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #1e3c50;
}
.welcome {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 60%;
  padding: 4rem;
  font-size: 2rem;
  color: #fafafa;
}
.intro {
  display: flex;
  flex-direction: row;
  height: 100%;
  padding: 2rem 1rem;
  justify-content: space-around;
}

.intro-img {
  width: 100%;
  height: 100%;
  border-radius: 2rem;
  object-fit: cover;
  border: 1px solid #798d8e;
}
.intro-det {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
  margin: 0 1rem;
}

.features {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 40%;
  padding: 4rem;
  justify-content: space-evenly;
}

.feature {
  display: flex;
  width: 100%;
  margin: 1rem 0;
  align-items: center;
}

.feature-icon {
  display: 'flex';
  min-height: 4rem;
  min-width: 4rem;
  border-radius: 1rem;
  background-color: #00cd6a;
}
.feature-det {
  display: 'flex';
  color: #fafafa;
  margin-left: 1rem;
}
</style>
