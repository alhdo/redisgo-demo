import { Router } from 'express'
import { StationService } from './stations.service'
import { BadRequestException, NotFoundException } from '~/utils/exceptions'

const StationsController = Router()

const service = new StationService()


StationsController.get('/', (req, res) => {
    return res.status(200).json('Data flow')
})


export { StationsController }