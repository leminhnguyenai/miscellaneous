import Worker from "./worker";

const workers = [
  new Worker(1),
  new Worker(2),
  new Worker(3),
  new Worker(4),
  new Worker(5),
  new Worker(6),
  new Worker(7),
  new Worker(8),
  new Worker(9),
  new Worker(10),
  new Worker(11),
  new Worker(12),
  new Worker(13),
  new Worker(14),
];

async function main() {
  for (let i = 0; i < workers.length; i++) {
    workers[i].setPeers(workers.slice(i));
  }

  for (let i = 0; i < workers.length; i++) {
    workers[i].run();
  }
}

main();
