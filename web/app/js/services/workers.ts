export interface WorkersCollection {
  fetchWorker: Worker;
}

const workers = {
  fetchWorker: new Worker("/fetch-worker.js")
};

export default workers;
