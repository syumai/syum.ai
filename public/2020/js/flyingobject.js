const FlyingObjectKinds = {
  EggPlant: 'EGG_PLANT',
  Hawk: 'HAWK',
};

const flyingObjectImagePaths = {
  [FlyingObjectKinds.EggPlant]: '/2020/images/EggPlant.svg',
  [FlyingObjectKinds.Hawk]: '/2020/images/Hawk.svg',
};

const flyingObjectClassNames = {
  [FlyingObjectKinds.EggPlant]: 'eggplant',
  [FlyingObjectKinds.Hawk]: 'hawk',
};

function getRandomKind() {
  const maxKinds = Object.values(FlyingObjectKinds).length;
  return Object.values(FlyingObjectKinds)[Math.floor(Math.random() * maxKinds)];
}

export class FlyingObject {
  constructor(container) {
    this.kind = null;
    this.top = 0;
    this.right = -150;
    this.transformX = 0;
    this.speed = 0;
    this.startedAt = null;
    this.el = document.createElement('img');
    this.render(); // Initial rendering

    this.container = container;
    container.appendChild(this.el);
  }

  get started() {
    return this.startedAt !== null;
  }

  destroy() {
    this.container.removeChild(this.el);
    this.el = null;
    this.container = null;
  }

  start(timestamp) {
    this.kind = getRandomKind();
    this.el.src = flyingObjectImagePaths[this.kind];
    this.el.className = 'flying-object';
    this.el.classList.add(flyingObjectClassNames[this.kind]);
    this.startedAt = timestamp;
    this.speed = 50 + Math.floor(Math.random() * 50);
    this.transformX = 0;
    this.top = Math.floor(Math.random() * 90);
  }

  update(timestamp) {
    if (this.transformX > this.container.clientWidth + 300) {
      this.startedAt = null;
      return;
    }
    this.calc(timestamp);
    this.render();
  }

  calc(timestamp) {
    const diff = timestamp - this.startedAt;
    const diffSeconds = diff / 1000;
    this.transformX = diffSeconds * this.speed;
  }

  render() {
    this.el.style.top = `${this.top}%`;
    this.el.style.right = `${this.right}px`;
    this.el.style.transform = `translateX(${-1 * this.transformX}px)`;
  }
}
