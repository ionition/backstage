/*
 * Copyright 2023 The Backstage Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { AnalyzeLocationEntityField } from '../models/AnalyzeLocationEntityField.model';
import { RecursivePartialEntity } from '../models/RecursivePartialEntity.model';

/**
 * This is some form of representation of what the analyzer could deduce. We should probably have a chat about how this can best be conveyed to the frontend. It'll probably contain a (possibly incomplete) entity, plus enough info for the frontend to know what form data to show to the user for overriding/completing the info.
 */
export interface AnalyzeLocationGenerateEntity {
  fields: Array<AnalyzeLocationEntityField>;
  entity: RecursivePartialEntity;
}
