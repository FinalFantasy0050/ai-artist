from googletrans import Translator
import logging

logger = logging.getLogger(__name__)

class PromptTranslator:
    def __init__(self):
        self.translator = Translator()
    
    def translate_prompt(self, prompt, src_lang='ko', dest_lang='en'):
        try:
            translated = self.translator.translate(prompt, src=src_lang, dest=dest_lang)
            translated_prompt = translated.text
            logger.info(f"* Translated prompt: {translated_prompt}")

            return translated_prompt
        except Exception as e:
            logger.error(f"Translation error: {e}")

            raise e
