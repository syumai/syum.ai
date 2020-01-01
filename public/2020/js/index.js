import { FlyingObject } from './flyingobject.js';

const flyingObjectContainer = document.getElementById('flyingObjectContainer');

const flyingObjects = [];

async function initializeFlyingObjects() {
  function addFlyingObject(timeout) {
    return new Promise(resolve => {
      window.setTimeout(() => {
        flyingObjects.push(new FlyingObject(flyingObjectContainer));
        resolve();
      }, timeout);
    });
  }
  for (let i = 0; i < 10; i++) {
    await addFlyingObject(800);
  }
}

function update(timestamp) {
  for (const obj of flyingObjects) {
    if (!obj.started) {
      obj.start(timestamp);
      continue;
    } else {
      obj.update(timestamp);
    }
  }
  window.requestAnimationFrame(update);
}

window.requestAnimationFrame(update);

window.addEventListener('DOMContentLoaded', () => {
  const mainContainer = document.getElementById('mainContainer');
  const sun = document.getElementById('sun');
  const ray = document.getElementById('ray');
  const overlay = document.getElementById('overlay');
  mainContainer.classList.remove('dark');
  mainContainer.classList.add('rising');
  function mainContainerTransition() {
    mainContainer.removeEventListener('transitionend', mainContainerTransition);
    mainContainer.classList.remove('rising');
    mainContainer.classList.add('day');
    initializeFlyingObjects();
  }
  mainContainer.addEventListener('transitionend', mainContainerTransition);
  sun.classList.remove('rising');
  sun.classList.add('risen');
  ray.classList.remove('rising');
  ray.classList.add('risen');
  overlay.classList.remove('hidden');
});
