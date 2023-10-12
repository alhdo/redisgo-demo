export enum JobStatus {
    Pending = 'pending',
    InProgress = 'in-progress',
    Completed = 'completed',
    Failed = 'failed',
  }
export enum TaskType {
    Fetch = 'fetch',
}
export interface Job {
    id: string;
    task: TaskType;
    params: Record<string, any>;
    data?: any;
    status: JobStatus
}