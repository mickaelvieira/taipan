import FastDomBase from "fastdom";
import fastdomPromised from "fastdom/extensions/fastdom-promised";

const FastDom = FastDomBase.extend(fastdomPromised);

export default class Scroll {
  // Header's height
  private marginTop = 70;

  // is the scroll active?
  private active = false;

  // previous scroll position
  private previous = 0;

  // current scroll position
  private position = 0;

  // record the scroll status
  record(active: boolean): void {
    this.active = active;
    if (active) {
      this.setPosition();
    }
  }

  async setPosition(): Promise<void> {
    const position = await FastDom.measure(() => window.scrollY);
    this.previous = this.position;
    this.position = position;
  }

  isActive(): boolean {
    return this.active;
  }

  getPosition(): number {
    return Math.abs(this.position) + this.marginTop;
  }

  isIdle(): boolean {
    return this.previous === this.position;
  }

  isUp(): boolean {
    return this.previous < this.position;
  }

  isDown(): boolean {
    return this.previous > this.position;
  }
}
