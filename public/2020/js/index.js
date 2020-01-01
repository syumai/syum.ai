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

initializeFlyingObjects();

window.requestAnimationFrame(update);
