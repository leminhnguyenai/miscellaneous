class Worker {
  private id: number;
  private peers: Worker[];
  public alive: boolean;
  public leader: boolean;
  constructor(id: number) {
    this.id = id;
    this.peers = [];
    //this.alive = true;
    this.alive = false;
    this.leader = false;
  }

  async run(): Promise<void> {
    while (this.alive) {
      const leader = this.peers.find((peer) => peer.isLeader());
      if (!leader)
        this.send(); // Inititate election if no leader is found
      else {
        if (!leader.alive) this.send(); // Initiate election if leader is down
        await new Promise((resolve) => setTimeout(resolve, 100));
      }
    }
  }

  getId(): number {
    return this.id;
  }

  setPeers(peers: Worker[]): void {
    this.alive = true;
    this.peers = peers;
  }

  becomeLeader(): void {
    console.log(`Node ${this.id} is the leader`);
    this.leader = true;
  }

  isLeader(): boolean {
    return this.leader;
  }

  async send(): Promise<void> {
    if (this.isLeader()) return;

    for (let i = 0; i < this.peers.length; i++) {
      if (this.peers[i].id > this.id) {
        console.log(
          `Node ${this.id} send message to node ${this.peers[i].getId()}`,
        );
        const res = this.peers[i].receive();
        if (res) {
          console.log(
            `Node ${this.peers[i].getId()} received message from  node ${this.id}`,
          );
          return;
        }
      }
    }

    this.becomeLeader();
  }

  receive(): boolean {
    if (this.alive) {
      this.send();
      return true;
    }
    return false;
  }
}

export default Worker;
