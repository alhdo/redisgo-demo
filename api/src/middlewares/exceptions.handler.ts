import { NextFunction, Request, Response } from 'express';

export const ExceptionHandler = (err: any, req: Request, res: Response, next: NextFunction) => {
  /**
   * Middleware to handle global error
   *
   * @param err - Express error or internal
   * @param req - Initial request
   * @param res - Response object
   * @param next - Go to the next middleware if available
   */
  if (res.headersSent) {
    return next(err);
  }

  if (err.status && err.error) {
    return res.status(err.status).json({ error: err.error });
  }

  return res.status(500).json({ error: 'Internal error' });
};
