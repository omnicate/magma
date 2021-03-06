/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

const express = require('express');

import type {ExpressRequest, ExpressResponse} from 'express';

const router: express.Router<
  ExpressRequest,
  ExpressResponse,
> = express.Router();

router.get('/', (req: ExpressRequest, res: ExpressResponse) => {
  res.status(200).end('Success');
});

module.exports = router;
