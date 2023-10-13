import { NotFoundException } from '~/utils/exceptions';

/**
 * For all undefined routes return an errror
 */

export const UnknownRoutesHandler = () => {
  throw new NotFoundException('The requested resources does not exist!');
};
